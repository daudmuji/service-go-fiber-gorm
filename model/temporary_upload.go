package model

import (
	"github.com/gofiber/fiber/v2"
	"golang-template-service/web/request"
	"golang-template-service/web/response"
)

type DataExcel struct {
	No          string
	Name        string
	PhoneNumber string
	Gender      string
	Address     string
}

type DataUpload struct {
	Name        string
	PhoneNumber string
	Gender      string
	Address     string
}

type SampleUpload struct {
	UUID                  string
	DataExcel             DataUpload `gorm:"embedded"`
	StatusValidation      bool
	InformationValidation string
}

type MappingDataUpload struct {
	Ctx          *fiber.Ctx
	Worker       int
	Start        int
	TotalData    int
	RawData      [][]string
	CountRowData int
	Req          *request.UploadRequest
}

type TemporaryUploadRepository interface {
	CreateDataUpload(sampleUpload *SampleUpload) (err error, result int64)
}

type TemporaryUploadUsecase interface {
	UploadBulkExcel(ctx *fiber.Ctx, req *request.UploadRequest) (err error, response response.WebResponse)
}
