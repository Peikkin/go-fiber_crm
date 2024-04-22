package main

import (
	"os"

	"github.com/Peikkin/go_fiber_crm/db"
	"github.com/Peikkin/go_fiber_crm/lead"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLeads)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func ConnectDB() {
	var err error
	dns := "host=localhost user=root password=1234 dbname=crm port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db.DBconn, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка подключения к базе данных")
	}
	log.Info().Msg("Подключение к базе данных успешно")

	err = db.DBconn.AutoMigrate(&lead.Lead{})
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка миграции базы данных")
	}
	log.Info().Msg("Миграция базы данных успешна")
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := fiber.New()
	ConnectDB()
	setupRoutes(app)
	log.Info().Msg("Запуск сервера на порту :8080")
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка запуска сервера")
	}
}
