package lead

import (
	"github.com/gofiber/fiber"
	"github.com/ushieru/go-fiber-crm/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(ctx *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	ctx.JSON(leads)
}

func GetLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	ctx.JSON(lead)
}

func CreateLead(ctx *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := ctx.BodyParser(lead); err != nil {
		ctx.Status(503).SendString("failed to parse body")
		return
	}
	db.Create(&lead)
	ctx.JSON(lead)
}

func DeleteLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		ctx.Status(500).Send("lead not found")
		return
	}
	db.Delete(&lead)
	ctx.Send("lead deleted")
}
