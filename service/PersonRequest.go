package service

type PersonRequest struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama" binding:"required"`
	Alamat string `json:"alamat" binding:"required"`
}
