package handlers

import (
	"fiber-gorm/middleware"
	"fiber-gorm/models"
	"fiber-gorm/repo"

	"github.com/gofiber/fiber/v2"
)

func GetReportA(ctx *fiber.Ctx) error {
	logedUser := middleware.GetLogedUser(ctx.Context())

	params := new(models.ReportRequest)
	if err := ctx.QueryParser(params); err != nil {
		return err
	}

	if params.Pagination == nil {
		params.Pagination = &models.Pagination{}
	}
	if params.Page == 0 {
		params.Page = models.DefaultPage
	}
	if params.Limit == 0 {
		params.Limit = models.DefaultLimit
	}

	data, err := repo.GetReportA(params, logedUser.ID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return ctx.JSON(data)
}

func GetReportB(ctx *fiber.Ctx) error {
	logedUser := middleware.GetLogedUser(ctx.Context())

	params := new(models.ReportRequest)
	if err := ctx.QueryParser(params); err != nil {
		return err
	}

	if params.Pagination == nil {
		params.Pagination = &models.Pagination{}
	}
	if params.Page == 0 {
		params.Pagination.Page = models.DefaultPage
	}
	if params.Limit == 0 {
		params.Pagination.Limit = models.DefaultLimit
	}

	data, err := repo.GetReportB(params, logedUser.ID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "data not found",
		})
	}

	return ctx.JSON(data)
}
