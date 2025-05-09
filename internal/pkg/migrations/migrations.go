package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

func RunMigrations(db *sql.DB) error {
	migrationDir, err := getMigrationsDir()
	if err != nil {
		return fmt.Errorf("erro ao localizar diretório de migrations: %v", err)
	}

	log.Printf("Usando diretório de migrations: %s", migrationDir)

	err = createMigrationsTable(db)
	if err != nil {
		return err
	}

	files, err := listMigrationFiles(migrationDir)
	if err != nil {
		return err
	}

	sort.Strings(files)

	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := filepath.Base(file)

		if appliedMigrations[fileName] {
			log.Printf("Migration '%s' já foi aplicada anteriormente", fileName)
			continue
		}

		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("erro ao ler arquivo de migration %s: %v", fileName, err)
		}

		log.Printf("Aplicando migration: %s", fileName)

		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("erro ao iniciar transação: %v", err)
		}

		_, err = tx.Exec(string(content))
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("erro ao executar migration %s: %v", fileName, err)
		}

		_, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", fileName)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("erro ao registrar migration %s: %v", fileName, err)
		}

		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("erro ao finalizar transação: %v", err)
		}

		log.Printf("Migração '%s' aplicada com sucesso", fileName)
	}

	log.Println("Todas as migrations foram aplicadas com sucesso")
	return nil
}

func getMigrationsDir() (string, error) {
	paths := []string{
		"internal/pkg/migrations",
		"../internal/pkg/migrations",
		"../../internal/pkg/migrations",
	}

	_, filename, _, ok := runtime.Caller(0)
	if ok {
		dir := filepath.Dir(filename)
		paths = append(paths, dir)
	}

	for _, path := range paths {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return path, nil
		}
	}

	return "", fmt.Errorf("diretório de migrations não encontrado")
}

func createMigrationsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version VARCHAR(255) NOT NULL PRIMARY KEY,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela de controle de migrations: %v", err)
	}

	return nil
}

func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	migrations := make(map[string]bool)

	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, fmt.Errorf("erro ao obter migrations aplicadas: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		migrations[version] = true
	}

	return migrations, nil
}

func listMigrationFiles(dir string) ([]string, error) {
	var files []string

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("diretório de migrations não encontrado: %s", dir)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	return files, nil
}
