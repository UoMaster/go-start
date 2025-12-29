package main

import (
	"log"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/database"
	"github.com/Caknoooo/go-gin-clean-starter/middlewares"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/providers"
	"github.com/samber/do"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	injector := do.New()

	providers.RegisterDependencies(injector)

	// Run database migration
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	if err := database.Migrate(db); err != nil {
		log.Fatalf("error migration: %v", err)
	}

	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	// Register module routes
	user.RegisterRoutes(server, injector)
	auth.RegisterRoutes(server, injector)

	run(server)
}

func run(server *gin.Engine) {
	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	log.Printf("Server running on %s", serve)

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
