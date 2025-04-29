package main

import (
	"log"

	_ "github.com/Felipe8297/go-contacts-api/docs"
	"github.com/Felipe8297/go-contacts-api/internal/contacts"
	"github.com/Felipe8297/go-contacts-api/internal/pkg/db"
	"github.com/Felipe8297/go-contacts-api/internal/pkg/migrations"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Inicializa a conexão com o banco de dados
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer database.Close()

	// Executa as migrações automaticamente na inicialização
	log.Println("Executando migrations pendentes...")
	err = migrations.RunMigrations(database)
	if err != nil {
		log.Fatalf("Erro ao executar migrações: %v", err)
	}

	// Configura o modo de produção (remove o warning de debug mode)
	gin.SetMode(gin.ReleaseMode)

	// Usa gin.New() em vez de gin.Default() para evitar o warning de middleware já anexado
	router := gin.New()

	// Adiciona manualmente os middlewares necessários
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Define uma lista confiável de proxies (resolve o warning de proxy)
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Inicialização das camadas da aplicação
	contactsRepo := contacts.NewPostgresRepository(database)
	contactsService := contacts.NewService(contactsRepo)
	contactsHandler := contacts.NewHandler(contactsService)

	// Registra as rotas
	contactsHandler.RegisterRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicia o servidor
	port := ":8080"
	log.Printf("Servidor iniciado na porta %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
