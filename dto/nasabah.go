package dto

import (
	"time"

	"github.com/google/uuid"
)

type NasabahCreateDto struct {
	ID        		 uuid.UUID  `gorm:"primary_key" json:"id" form:"id"`
	Nama             string 	`json:"nama" form:"nama" binding:"required"`
	NoKtp            string 	`json:"no_ktp" form:"no_ktp" binding:"required"`
	TempatLahir      string 	`json:"tempat_lahir" form:"tempat_lahir" binding:"required"`
	TanggalLahir     time.Time 	`json:"tanggal_lahir" form:"tanggal_lahir" binding:"required" time_format:"2006-01-02"`
	StatusKawin      string 	`json:"status_kawin" form:"status_kawin" binding:"required"`
	Alamat           string 	`json:"alamat" form:"alamat" binding:"required"`
	Pekerjaan        string 	`json:"pekerjaan" form:"pekerjaan" binding:"required"`
	NamaPerusahaan   string 	`json:"nama_perusahaan" form:"nama_perusahaan" binding:"required"`
	AlamatPerusahaan string 	`json:"alamat_perusahaan" form:"alamat_perusahaan" binding:"required"`
	NoHp             string 	`json:"no_hp" form:"no_hp" binding:"required"`
	Email            string 	`json:"email" form:"email" binding:"required,email"`
	NamaIbuKandung   string 	`json:"nama_ibu_kandung" form:"nama_ibu_kandung" binding:"required"`
	JenisKelamin     string 	`json:"jenis_kelamin" form:"jenis_kelamin" binding:"required"`
}

type NasabahUpdateDto struct {
	ID        		 uuid.UUID  `gorm:"primary_key" json:"id" form:"id"`
	Nama             string 	`json:"nama" form:"nama"`
	NoKtp            string 	`json:"no_ktp" form:"no_ktp"`
	TempatLahir      string 	`json:"tempat_lahir" form:"tempat_lahir"`
	TanggalLahir     time.Time 	`json:"tanggal_lahir" form:"tanggal_lahir" time_format:"2006-01-02"`
	StatusKawin      string 	`json:"status_kawin" form:"status_kawin"`
	Alamat           string 	`json:"alamat" form:"alamat"`
	Pekerjaan        string 	`json:"pekerjaan" form:"pekerjaan"`
	NamaPerusahaan   string 	`json:"nama_perusahaan" form:"nama_perusahaan"`
	AlamatPerusahaan string 	`json:"alamat_perusahaan" form:"alamat_perusahaan"`
	NoHp             string 	`json:"no_hp" form:"no_hp"`
	Email            string 	`json:"email" form:"email"`
	NamaIbuKandung   string 	`json:"nama_ibu_kandung" form:"nama_ibu_kandung"`
	JenisKelamin     string 	`json:"jenis_kelamin" form:"jenis_kelamin"`
}