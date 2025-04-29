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

// RunMigrations executa todas as migrations SQL na pasta de migrations
func RunMigrations(db *sql.DB) error {
	// Define o diretório onde estão os arquivos de migração
	migrationDir, err := getMigrationsDir()
	if err != nil {
		return fmt.Errorf("erro ao localizar diretório de migrations: %v", err)
	}

	log.Printf("Usando diretório de migrations: %s", migrationDir)

	// Verifica se a tabela de migrations existe, caso contrário, cria
	err = createMigrationsTable(db)
	if err != nil {
		return err
	}

	// Lista todos os arquivos SQL no diretório
	files, err := listMigrationFiles(migrationDir)
	if err != nil {
		return err
	}

	// Ordena os arquivos pelo nome para executar na ordem correta
	sort.Strings(files)

	// Obtém migrations já executadas
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		return err
	}

	// Executa cada migration que ainda não foi aplicada
	for _, file := range files {
		fileName := filepath.Base(file)

		// Verifica se a migration já foi aplicada
		if appliedMigrations[fileName] {
			log.Printf("Migration '%s' já foi aplicada anteriormente", fileName)
			continue
		}

		// Lê o conteúdo do arquivo SQL
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("erro ao ler arquivo de migration %s: %v", fileName, err)
		}

		// Executa a migration
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

		// Registra que a migration foi executada
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

// getMigrationsDir tenta encontrar o diretório de migrations de várias maneiras
func getMigrationsDir() (string, error) {
	// Tenta encontrar o diretório a partir do diretório atual
	paths := []string{
		"internal/pkg/migrations",
		"../internal/pkg/migrations",
		"../../internal/pkg/migrations",
	}

	// Adiciona outra tentativa baseada no arquivo em execução
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		// Obtém o diretório do arquivo atual (migrations.go)
		dir := filepath.Dir(filename)
		paths = append(paths, dir)
	}

	// Tenta cada caminho
	for _, path := range paths {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return path, nil
		}
	}

	// Se nenhum caminho funcionou, retorna erro
	return "", fmt.Errorf("diretório de migrations não encontrado")
}

// createMigrationsTable cria a tabela que controla as migrations executadas
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

// getAppliedMigrations retorna um mapa de migrations já aplicadas
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

// listMigrationFiles lista todos os arquivos SQL no diretório de migrations
func listMigrationFiles(dir string) ([]string, error) {
	var files []string

	// Verifica se o diretório existe
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("diretório de migrations não encontrado: %s", dir)
	}

	// Lista os arquivos do diretório
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Filtra apenas os arquivos SQL
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	return files, nil
}
