package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang-template-service/model"
	"golang-template-service/web/request"
	"golang-template-service/web/response"
	"net/http"
	"time"
)

type SampleUploadController struct {
	SampleUpload model.TemporaryUploadUsecase
}

func NewSampleUploadController(r fiber.Router, SampleUpload model.TemporaryUploadUsecase) {
	handler := &SampleUploadController{SampleUpload}
	r.Post("sample-upload", handler.BulkUpload)
}

func (s SampleUploadController) BulkUpload(c *fiber.Ctx) error {

	req := new(request.UploadRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "invalid request",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	if err := validator.New().Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.WebResponse{
			Message:   "invalid request",
			TimeStamp: time.Now(),
			Data:      nil,
			Error:     err.Error(),
		})
	}

	err, webResponse := s.SampleUpload.UploadBulkExcel(c, req)
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
