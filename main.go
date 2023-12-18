package main

import (
	"gofr.dev/cmd/gofr/migration"
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
	"gofr.dev/pkg/gofr"

	"sample/handler"
	"sample/migrations"
	"sample/store"
)

func main() {
	// Creating GoFr app
	app := gofr.New()

	// Running migrations - UP
	if err := migration.Migrate("remote-config-data", dbmigration.NewGorm(app.GORM()),
		migrations.All(), dbmigration.UP, app.Logger); err != nil {
		app.Logger.Fatalf("Error in running migrations: %v", err)
	}

	productStore := store.New()
	productHandler := handler.New(productStore)

	// Creating routes
	app.POST("/product", productHandler.Create)
	app.GET("/products", productHandler.GetAll)
	app.GET("/product/{id}", productHandler.GetByID)
	app.PUT("/product/{id}", productHandler.Update)
	app.DELETE("/product/{id}", productHandler.Delete)

	// Starting server
	app.Start()
}
