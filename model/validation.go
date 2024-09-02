package model

type ValidationUsecase interface {
	FieldValidation(dataUpload DataUpload) []byte
}
