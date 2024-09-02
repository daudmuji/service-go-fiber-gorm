package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang-template-service/model"
	"golang-template-service/web/request"
	"golang-template-service/web/response"
	"net/http"
	"strconv"
	"time"
)

type ProductController struct {
	ProductUsecase model.ProductUsecase
}

func NewProductController(r fiber.Router, productUsecase model.ProductUsecase) {
	handler := &ProductController{productUsecase}
	r.Post("/product", handler.Create)
	r.Get("/product/list", handler.GetList)
	r.Get("/product/:nomorSeri", handler.GetDetail)
	r.Put("/product/:nomorSeri", handler.Update)
	r.Delete("/product/:nomorSeri", handler.Delete)
}

func (p ProductController) Create(c *fiber.Ctx) error {

	req := new(request.ProductCreateRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "invalid request for create product",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	if err := validator.New().Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "validator invalid request for create product",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	err, webResponse := p.ProductUsecase.CreateProduct(c, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "internal error",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(webResponse)
}

func (p ProductController) GetList(c *fiber.Ctx) error {

	err, successResponse := p.ProductUsecase.GetListProduct(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "internal error get list data",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(successResponse)
}

func (p ProductController) GetDetail(c *fiber.Ctx) error {

	query := c.Params("nomorSeri", "0")
	atoi, err := strconv.Atoi(query)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "error parse string to int",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	err, successResponse := p.ProductUsecase.GetDetailProductByNomorSeri(c, atoi)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "internal error get list data",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(successResponse)
}

func (p ProductController) Update(c *fiber.Ctx) error {

	query := c.Params("nomorSeri", "0")
	atoi, err := strconv.Atoi(query)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "error parse string to int",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	req := new(request.ProductUpdateRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "invalid request for update",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	if err := validator.New().Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "validator invalid request for update",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	err, webResponse := p.ProductUsecase.UpdateProductByNomorSeri(c, atoi, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "internal error",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(webResponse)
}

func (p ProductController) Delete(c *fiber.Ctx) error {

	query := c.Params("nomorSeri", "0")
	atoi, err := strconv.Atoi(query)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "error parse string to int",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	err, successResponse := p.ProductUsecase.DeleteProductByNomorSeri(c, atoi)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "internal error",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(successResponse)
}
