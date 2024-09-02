package request

type ProductCreateRequest struct {
	NamaBarang       string            `json:"nama_barang" validate:"required"`
	JumlahStokBarang int               `json:"jumlah_stok_barang" validate:"required"`
	NomorSeriBarang  int               `json:"nomor_seri_barang" validate:"required"`
	AdditionalInfo   map[string]string `json:"additional_info" validate:"required"`
	GambarBarang     string            `json:"gambar_barang" validate:"required"`
}

type ProductUpdateRequest struct {
	NamaBarang       string      `json:"nama_barang"`
	JumlahStokBarang int         `json:"jumlah_stok_barang"`
	NomorSeriBarang  int         `json:"nomor_seri_barang"`
	AdditionalInfo   interface{} `json:"additional_info"`
	GambarBarang     string      `json:"gambar_barang"`
}
