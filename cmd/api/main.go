package main

import (
	"log"

	_ "github.com/Felipe8297/go-contacts-api/docs"
	"github.com/Felipe8297/go-contacts-api/internal/contacts"
	"github.com/Felipe8297/go-contacts-api/internal/pkg/db"
	"github.com/Felipe8297/go-contacts-api/internal/pkg/middleware"
	"github.com/Felipe8297/go-contacts-api/internal/pkg/migrations"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer database.Close()

	log.Println("Executando migrations pendentes...")
	err = migrations.RunMigrations(database)
	if err != nil {
		log.Fatalf("Erro ao executar migrações: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.MetricsMiddleware())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	contactsRepo := contacts.NewPostgresRepository(database)
	contactsService := contacts.NewService(contactsRepo)
	contactsHandler := contacts.NewHandler(contactsService)

	contactsHandler.RegisterRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoint para métricas do Prometheus
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	port := "0.0.0.0:8080"
	log.Printf("Servidor iniciado na porta %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
