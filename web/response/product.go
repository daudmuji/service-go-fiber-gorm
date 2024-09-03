package response

type ProductResponse struct {
	ID               string
	NamaBarang       string
	JumlahStokBarang string
	NomorSeriBarang  int64
	AdditionalInfo   map[string]string
	GambarBarang     string
	CreatedAt        string
	UpdatedAt        string
}
