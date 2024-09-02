package usecase

import (
	"encoding/json"
	"fmt"
	"golang-template-service/model"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type ValidationUsecase struct {
}

func NewValidationUsecase() *ValidationUsecase {
	return &ValidationUsecase{}
}

var ValidationUpload = map[string]string{
	"name":         "M",
	"phone_number": "M,min:10",
	"gender":       "M,char:A",
}

func (u *ValidationUsecase) FieldValidation(dataUpload model.DataUpload) []byte {

	informationValidation := map[string]string{}
	var excelFieldValidation []string
	var resultValidation string

	excelFieldValidation = strings.Split(ValidationUpload["name"], ",")
	resultValidation = u.LoopValidation(dataUpload.Name, excelFieldValidation)
	if resultValidation != "" {
		informationValidation["name"] = resultValidation
	}

	excelFieldValidation = strings.Split(ValidationUpload["phone_number"], ",")
	resultValidation = u.LoopValidation(dataUpload.PhoneNumber, excelFieldValidation)
	if resultValidation != "" {
		informationValidation["phone_number"] = resultValidation
	}

	excelFieldValidation = strings.Split(ValidationUpload["gender"], ",")
	resultValidation = u.LoopValidation(dataUpload.Gender, excelFieldValidation)
	if resultValidation != "" {
		informationValidation["gender"] = resultValidation
	}

	validationByte, err := json.Marshal(informationValidation)
	if err != nil {
		e := fmt.Sprintf("Error marshalling JSON Validation | Error : %s", err.Error())
		log.Println(e)
		return validationByte
	}
	return validationByte
}

func (u *ValidationUsecase) LoopValidation(data string, value []string) string {
	var description string

	for _, v := range value {
		resultValidation := u.NewValidation(v, data)
		if resultValidation != "" {
			if description == "" {
				description = resultValidation
			} else {
				description += " || " + resultValidation
			}
		}
	}
	return description
}

func (u *ValidationUsecase) NewValidation(v, data string) string {

	if v == "M" {
		if !MandatoryValidation(data) {
			return "Harus di isi"
		}
	} else if strings.Contains(v, "min") {
		minLength, _ := strconv.Atoi(strings.Split(v, ":")[1])
		if !MinLengthValidation(data, minLength) {
			return fmt.Sprintf("Minimal memiliki %d karakter", minLength)
		}
	} else if strings.Contains(v, "char") {
		ref := strings.Split(v, ":")[1]
		if !CharValidation(data, ref) {
			switch ref {
			case "A":
				return "Input Hanya Boleh Alfabet"
			case "N":
				return "Input Hanya Boleh Angka"
			case "AN":
				return "Input Bisa Alfabet dan Angka"
			}
		}
	}

	return ""
}

func MandatoryValidation(data string) bool {
	if data == "" {
		return false
	}

	return true
}

func MinLengthValidation(data string, minLength int) bool {
	if data == "" {
		return false
	} else if len(data) < minLength {
		return false
	}
	return true
}

func CharValidation(data, ref string) bool {
	if data == "" {
		return true
	}

	switch ref {
	case "A":
		return regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(data)
	case "N":
		return regexp.MustCompile(`^[0-9]*$`).MatchString(data)
	case "AN":
		return regexp.MustCompile(`^[a-zA-Z0-9\s]*$`).MatchString(data)
	}
	return false
}
