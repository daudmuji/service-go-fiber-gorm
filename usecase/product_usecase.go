package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-template-service/model"
	"golang-template-service/web/request"
	"golang-template-service/web/response"
	"log"
	"time"
)

type ProductUsecase struct {
	ProductRepo model.ProductRepository
}

func NewProductUsecase(productRepo model.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo}
}

func (p *ProductUsecase) CreateProduct(ctx *fiber.Ctx, request *request.ProductCreateRequest) (err error, res response.SuccessResponse) {

	marshal, err := json.Marshal(request.AdditionalInfo)
	if err != nil {
		e := fmt.Sprintf("Error Marshal Json On Create : %s, nomor seri product: %d | Error: %s", request.AdditionalInfo, request.NomorSeriBarang, err.Error())
		log.Println(e)
		return err, res
	}

	product := model.Product{
		NamaBarang:       request.NamaBarang,
		JumlahStokBarang: request.JumlahStokBarang,
		NomorSeriBarang:  request.NomorSeriBarang,
		AdditionalInfo:   string(marshal),
		GambarBarang:     request.GambarBarang,
		CreatedAt:        time.Now(),
	}

	err, result := p.ProductRepo.CreateProduct(&product)
	if err != nil {
		e := fmt.Sprintf("Error inserting product with nama: %s, nomor seri: %d | Error: %s", request.NamaBarang, request.NomorSeriBarang, err.Error())
		log.Println(e)
		return err, res
	}

	if result == 0 {
		e := fmt.Sprintf("Inserting data not rows affected : %s, nomor seri: %d | Error: %s", request.NamaBarang, request.NomorSeriBarang)
		log.Println(e)
		return err, res
	}

	res = response.SuccessResponse{
		Message:   "Success Store",
		TimeStamp: time.Now(),
		Data:      request,
	}

	return nil, res

}

func (p *ProductUsecase) GetListProduct(ctx *fiber.Ctx) (err error, res response.SuccessResponse) {

	err, products := p.ProductRepo.GetListProducts()
	if err != nil {
		e := fmt.Sprintf("Error getting products: %s", err.Error())
		log.Println(e)
		return err, res
	}

	if len(products) == 0 {
		e := fmt.Sprintf("No products found")
		log.Println(e)
		return err, res
	}

	return nil, response.SuccessResponse{
		Message:   "Success Get List Product",
		TimeStamp: time.Now(),
		Data:      products,
	}
}

func (p *ProductUsecase) GetDetailProductByNomorSeri(ctx *fiber.Ctx, nomorSeri int) (err error, res response.SuccessResponse) {

	err, product := p.ProductRepo.GetDetailProductByNomorSeri(nomorSeri)
	if err != nil {
		e := fmt.Sprintf("Error getting products: %s", err.Error())
		log.Println(e)
		return err, res
	}

	if product.NomorSeriBarang.Int64 == 0 {
		e := fmt.Sprintf("No products found")
		log.Println(e)
		return err, res
	}

	return nil, response.SuccessResponse{
		Message:   "Success Get List Product",
		TimeStamp: time.Now(),
		Data:      product,
	}
}

func (p *ProductUsecase) UpdateProductByNomorSeri(ctx *fiber.Ctx, nomorSeri int, request *request.ProductUpdateRequest) (err error, res response.SuccessResponse) {

	marshal, err := json.Marshal(request.AdditionalInfo)
	if err != nil {
		e := fmt.Sprintf("Error Marshal Json On Update : %s, nomor seri product: %d | Error: %s", request.AdditionalInfo, request.NomorSeriBarang, err.Error())
		log.Println(e)
		return err, res
	}

	product := model.Product{
		NamaBarang:       request.NamaBarang,
		JumlahStokBarang: request.JumlahStokBarang,
		NomorSeriBarang:  request.NomorSeriBarang,
		AdditionalInfo:   string(marshal),
		GambarBarang:     request.GambarBarang,
		UpdatedAt:        time.Now(),
	}

	err, result := p.ProductRepo.UpdateProductByNomorSeri(nomorSeri, &product)
	if err != nil {
		e := fmt.Sprintf("Error update product with nama: %s, nomor seri: %d | Error: %s", request.NamaBarang, request.NomorSeriBarang, err.Error())
		log.Println(e)
		return err, res
	}

	if result == 0 {
		e := fmt.Sprintf("Updating data not rows affected : %s, nomor seri: %d | Error: %s", request.NamaBarang, request.NomorSeriBarang)
		log.Println(e)
		return err, res
	}

	res = response.SuccessResponse{
		Message:   "Success Update Product",
		TimeStamp: time.Now(),
		Data:      request,
	}

	return nil, res
}

func (p *ProductUsecase) DeleteProductByNomorSeri(ctx *fiber.Ctx, nomorSeri int) (err error, res response.SuccessResponse) {

	var product *model.Product

	err, _ = p.ProductRepo.DeleteProductByNomorSeri(nomorSeri, product)
	if err != nil {
		e := fmt.Sprintf("Error deleting product with no seri: %d, Error: %s", nomorSeri, err.Error())
		log.Println(e)
		return err, res
	}

	res = response.SuccessResponse{
		Message:   "Success Delete",
		TimeStamp: time.Now(),
		Data:      nil,
	}

	return nil, res
}
