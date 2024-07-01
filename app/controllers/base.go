package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	dto "main.go/app/dtos"
	"main.go/app/repositories"
	"main.go/app/services"
	"main.go/pkg/middlewares"
	"main.go/pkg/response"
	"main.go/platform/database"
)

// BaseCreate func for creates a new base.
// @Description Create a new base.
// @Summary create a new base
// @Tags Base
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param age body string true "Age"
// @Success 201 {string} status "created"
// @Router /api/v1/base [post]
func BaseCreate(ctx *fiber.Ctx) error {
	repo := repositories.NewIKBaseRepository(database.DB)
	service := services.NewIkBaseService(repo)
	var request dto.BaseRequest
	err := middlewares.ParseAndValidate(ctx, &request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ResponseData(nil, nil, err, fiber.StatusUnprocessableEntity))
	}
	err = service.Create(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ResponseData(nil, nil, err, fiber.StatusBadRequest))
	}
	return ctx.Status(fiber.StatusCreated).JSON(response.ResponseData(nil, nil, nil, fiber.StatusCreated))
}

// UpdateBase func for updates base by given ID.
// @Description Update base.
// @Summary update base
// @Tags Base
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param age body string true "Age"
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/base [put]
func BaseUpdate(ctx *fiber.Ctx) error {
	repo := repositories.NewIKBaseRepository(database.DB)
	service := services.NewIkBaseService(repo)
	id := ctx.Params("id")
	var request dto.BaseRequest
	err := middlewares.ParseAndValidate(ctx, &request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ResponseData(nil, nil, err, fiber.StatusUnprocessableEntity))
	}
	err = service.Update(id, &request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ResponseData(nil, nil, err, fiber.StatusBadRequest))
	}
	return ctx.Status(fiber.StatusOK).JSON(response.ResponseData(nil, nil, nil, fiber.StatusOK))

}

// GetBase func gets all exists base.
// @Description Get all exists books.
// @Summary get all exists base
// @Tags Base
// @Accept json
// @Produce json
// @Success 200 {array} dto.Response
// @Router /api/v1/base [get]
func BaseList(ctx *fiber.Ctx) error {
	repo := repositories.NewIKBaseRepository(database.DB)
	service := services.NewIkBaseService(repo)
	results, count, err := service.GetList(ctx.Query("limit"), ctx.Query("skip"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ResponseData(nil, nil, err, fiber.StatusBadRequest))
	}
	metaDataResponse, err := response.ResponseMetaData(int(count), ctx.Query("limit"), ctx.Query("skip"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ResponseData(nil, nil, err, fiber.StatusBadRequest))

	}
	return ctx.Status(fiber.StatusOK).JSON(response.ResponseData(results, &metaDataResponse, nil, fiber.StatusOK))

}

// GetBook func gets base by given ID or 404 error.
// @Description Get base by given ID.
// @Summary get base by given ID
// @Tags Base
// @Accept json
// @Produce json
// @Param id path string true "Base ID"
// @Success 200 {object} dto.Response
// @Router /api/v1/base/{id} [get]
func BaseGetById(ctx *fiber.Ctx) error {
	repo := repositories.NewIKBaseRepository(database.DB)
	service := services.NewIkBaseService(repo)
	id := ctx.Params("id")
	result, err := service.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ResponseData(nil, nil, err, fiber.StatusNotFound))
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ResponseData(nil, nil, err, fiber.StatusBadRequest))
	}
	return ctx.Status(fiber.StatusOK).JSON(response.ResponseData(result, nil, nil, fiber.StatusOK))

}

// BaseDelete func for deletes base by given ID.
// @Description Delete base by given ID.
// @Summary delete base by given ID
// @Tags Base
// @Accept json
// @Produce json
// @Param id body string true "Base ID"
// @Success 200 {string} status "ok"
// @Router /v1/base [delete]
func BaseDelete(ctx *fiber.Ctx) error {
	repo := repositories.NewIKBaseRepository(database.DB)
	service := services.NewIkBaseService(repo)
	id := ctx.Params("id")
	err := service.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ResponseData(nil, nil, err, fiber.StatusBadRequest))

	}
	return ctx.Status(fiber.StatusOK).JSON(response.ResponseData(nil, nil, nil, fiber.StatusOK))

}
