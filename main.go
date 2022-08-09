package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/ushieru/go-fiber-crm/database"
	"github.com/ushieru/go-fiber-crm/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupLeadRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Print("database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupLeadRoutes(app)
	app.Listen(3000)
}
