package controllers

import (
	"project-management/models"
	"project-management/services"
	"project-management/utils"

	"github.com/gofiber/fiber/v2"
)

type ListController struct {
	service services.ListService
}

func NewListController(s services.ListService) *ListController {
	return &ListController{
		service: s,
	}
}

func (c *ListController) CreateList(ctx *fiber.Ctx) error {
	list := new(models.List)

	if err := ctx.BodyParser(list); err != nil {
		return utils.BadRequest(ctx, "invalid request body", err.Error())
	}

	if err := c.service.Create(list); err != nil {
		return utils.BadRequest(ctx, "Gagal membuat list", err.Error())
	}

	return utils.Success(ctx, "List berhasil dibuat", list)
}
