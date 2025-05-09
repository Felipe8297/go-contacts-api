package main

import (
	"log"

	"github.com/Felipe8297/go-contacts-api/internal/pkg/db"
	"github.com/Felipe8297/go-contacts-api/internal/pkg/migrations"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	log.Println("Iniciando a execução das migrations...")
	err = migrations.RunMigrations(db)
	if err != nil {
		log.Fatalf("Erro ao executar migrations: %v", err)
	}

	log.Println("Migrações concluídas com sucesso!")
}
