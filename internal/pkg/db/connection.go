package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	host := "postgres"
	port := "5432"
	user := "docker"
	password := "docker"
	dbname := "contactsdb"

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir conexão: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("falha ao pingar o banco de dados: %v", err)
	}

	log.Println("Conexão com PostgreSQL estabelecida com sucesso!")
	return db, nil
}
