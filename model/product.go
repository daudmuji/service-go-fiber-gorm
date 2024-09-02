package model

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"golang-template-service/web/request"
	"golang-template-service/web/response"
	"time"
)

type Product struct {
	NamaBarang       string
	JumlahStokBarang int
	NomorSeriBarang  int
	AdditionalInfo   string
	GambarBarang     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ProductQuery struct {
	ID               sql.NullString
	NamaBarang       sql.NullString
	JumlahStokBarang sql.NullString
	NomorSeriBarang  sql.NullInt64
	AdditionalInfo   sql.NullString
	GambarBarang     sql.NullString
	CreatedAt        sql.NullString
	UpdatedAt        sql.NullString
}

type ProductRepository interface {
	CreateProduct(product *Product) (err error, result int64)
	GetListProducts() (err error, products []ProductQuery)
	GetDetailProductByNomorSeri(nomorSeri int) (err error, product ProductQuery)
	UpdateProductByNomorSeri(nomorSeri int, product *Product) (err error, result int64)
	DeleteProductByNomorSeri(nomorSeri int, product *Product) (err error, result int64)
}

type ProductUsecase interface {
	CreateProduct(ctx *fiber.Ctx, productRequest *request.ProductCreateRequest) (err error, response response.SuccessResponse)
	GetListProduct(ctx *fiber.Ctx) (err error, response response.SuccessResponse)
	GetDetailProductByNomorSeri(ctx *fiber.Ctx, nomorSeri int) (err error, response response.SuccessResponse)
	UpdateProductByNomorSeri(ctx *fiber.Ctx, nomorSeri int, productRequest *request.ProductUpdateRequest) (err error, result response.SuccessResponse)
	DeleteProductByNomorSeri(ctx *fiber.Ctx, nomorSeri int) (err error, result response.SuccessResponse)
}
