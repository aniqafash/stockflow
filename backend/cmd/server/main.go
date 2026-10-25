package main

import (
	"context"
	"log"
	"stockflow/internal/handler"
	"stockflow/internal/pkg/config"
	"stockflow/internal/pkg/database"
	"stockflow/internal/repository"
	"stockflow/internal/route"
	"stockflow/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Open(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	r := route.NewRouter(route.Handlers{
		Product: productHandler,
	})

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
