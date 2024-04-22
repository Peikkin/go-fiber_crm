package lead

import (
	"github.com/Peikkin/go_fiber_crm/db"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(ctx *fiber.Ctx) error {
	db := db.DBconn
	var leads []Lead
	db.Find(&leads)
	ctx.JSON(leads)
	return nil
}

func GetLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := db.DBconn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		log.Error().Msg("Записи не найдены")
		return ctx.Status(500).SendString("Записи не найдены")
	}
	db.Find(&lead, id)
	ctx.JSON(lead)
	return nil
}

func NewLeads(ctx *fiber.Ctx) error {
	db := db.DBconn
	lead := new(Lead)
	if err := ctx.BodyParser(lead); err != nil {
		log.Error().Err(err).Msg("Ошибка создания клиента")
		return ctx.Status(503).SendString(err.Error())
	}
	db.Create(&lead)
	ctx.JSON(lead)
	return nil
}

func DeleteLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := db.DBconn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		log.Error().Msg("Записи не найдены")
		return ctx.Status(500).SendString("Записи не найдены")
	}
	db.Delete(&lead)
	ctx.SendString("Клиент успешно удален")
	ctx.JSON(lead)
	return nil
}
