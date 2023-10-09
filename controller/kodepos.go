package controller

import (
	"go-api/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type KodePosController struct {
	client *http.Client
}

func NewKodePosController(client *http.Client) KodePosController {
	return KodePosController{client: client}
}

func (kc KodePosController) GetAllKodePos(c *fiber.Ctx) error {
	kodePosService := service.NewKodePosService(kc.client)
	result, err := kodePosService.GetAllKodePos()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Berhasil memanggil list kode pos",
		"data":    result,
	})
}
