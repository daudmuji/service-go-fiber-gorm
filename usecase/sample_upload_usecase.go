package usecase

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"golang-template-service/helpers"
	"golang-template-service/model"
	"golang-template-service/web/request"
	"golang-template-service/web/response"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type SampleUploadUsecase struct {
	SampleUploadRepo  model.TemporaryUploadRepository
	ValidationUsecase model.ValidationUsecase
}

func NewSampleUploadUsecase(SampleUploadRepo model.TemporaryUploadRepository, ValidationUsecase model.ValidationUsecase) *SampleUploadUsecase {
	return &SampleUploadUsecase{SampleUploadRepo: SampleUploadRepo, ValidationUsecase: ValidationUsecase}
}

func (u *SampleUploadUsecase) UploadBulkExcel(ctx *fiber.Ctx, req *request.UploadRequest) (err error, res response.WebResponse) {

	res = response.WebResponse{
		Message:   "",
		TimeStamp: time.Now(),
		Data:      nil,
		Error:     nil,
	}

	fileB64 := req.Content
	decodeStr, err := base64.StdEncoding.DecodeString(fileB64)
	if err != nil {
		res.Message = "Error decode string"
		res.Error = err.Error()
		e := fmt.Sprintf("Error decode string err: %s", err.Error())
		log.Println(e)
		return err, res
	}

	file := bytes.NewReader(decodeStr)
	rawdata, countRowData, err := u.readFileExcel(file)
	if err != nil {
		res.Message = "Error read file"
		res.Error = err.Error()
		e := fmt.Sprintf("Error read file err: %s", err.Error())
		log.Println(e)
		return err, res
	}

	var mappingData model.MappingDataUpload
	var wg sync.WaitGroup

	limit, err := strconv.Atoi(os.Getenv("LIMIT_DATA_UPLOAD"))
	if err != nil {
		res.Message = "Error parse limit"
		res.Error = err.Error()
		e := fmt.Sprintf("Error convert limit data upload from env err: %s", err.Error())
		log.Println(e)
		return err, res
	}

	totalLoop := int(math.Ceil(float64(countRowData) / float64(limit)))

	start := 0
	for i := 0; i < totalLoop; i++ {
		if i == 0 {
			start = (i * limit) + 1
		} else {
			start = (i * limit) + 1
		}
		totalData := (i + 1) * limit
		worker := i + 1

		if totalData > countRowData {
			totalData = (countRowData - 1)
		}

		wg.Add(1)

		mappingData = model.MappingDataUpload{
			Ctx:          ctx,
			Worker:       worker,
			Start:        start,
			TotalData:    totalData,
			RawData:      rawdata,
			CountRowData: countRowData,
			Req:          req,
		}

		go func(mappingData model.MappingDataUpload) {
			defer wg.Done()
			u.mappingDataUpload(mappingData)
		}(mappingData)
	}

	res.Message = "Success Uplad Data"
	res.Data = fmt.Sprintf("Total Data %d", (len(rawdata) - 1))
	return nil, res
}

func (u *SampleUploadUsecase) readFileExcel(file *bytes.Reader) (rawdata [][]string, countRowData int, err error) {
	reader, err := excelize.OpenReader(file)
	if err != nil {
		e := fmt.Sprintf("open reader error: %s", err.Error())
		log.Println(e)
		return nil, countRowData, err
	}

	rawdata, err = reader.GetRows(reader.GetSheetName(reader.GetActiveSheetIndex()))
	if err != nil {
		e := fmt.Sprintf("get rows error: %s", err.Error())
		log.Println(e)
		return nil, countRowData, err
	}

	for i := 0; i < len(rawdata); i++ {
		rows := rawdata[i]
		if rows != nil {
			row := rows[1]
			if row != "" {
				countRowData++
			}
		}
	}
	return rawdata, countRowData, nil
}

func (u *SampleUploadUsecase) mappingDataUpload(mappingData model.MappingDataUpload) {

	var dataExcel model.DataExcel
	var dataUpload model.DataUpload
	var sampleUpload model.SampleUpload
	var isValid = true

	for i := mappingData.Start; i < mappingData.TotalData+1; i++ {

		if i == len(mappingData.RawData) {
			break
		}

		row := mappingData.RawData[i]

		row = helpers.AppendIfLess(row, 5)

		dataExcel = model.DataExcel{}

		v := reflect.ValueOf(&dataExcel).Elem()
		for j := 0; j < 5; j++ {
			v2 := v.Field(j)
			if row[j] != "" {
				v2.Set(reflect.ValueOf(row[j]))
			}
		}

		dataUpload = model.DataUpload{
			Name:        dataExcel.Name,
			PhoneNumber: dataExcel.PhoneNumber,
			Gender:      dataExcel.Gender,
			Address:     dataExcel.Address,
		}

		validationByte := u.ValidationUsecase.FieldValidation(dataUpload)
		if len(validationByte) > 2 {
			isValid = false
		}

		sampleUpload = model.SampleUpload{
			UUID:                  mappingData.Req.UUID,
			DataExcel:             dataUpload,
			StatusValidation:      isValid,
			InformationValidation: string(validationByte),
		}

		err, _ := u.SampleUploadRepo.CreateDataUpload(&sampleUpload)
		if err != nil {
			e := fmt.Sprintf("Error CreateDataUpload: %s", err.Error())
			log.Println(e)
		}
	}
}
