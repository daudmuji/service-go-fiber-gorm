package request

type UploadRequest struct {
	UUID    string `json:"uuid" validate:"required"`
	Content string `json:"content" validate:"required"`
}
