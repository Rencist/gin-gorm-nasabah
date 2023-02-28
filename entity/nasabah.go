package entity

import (
	"time"

	"github.com/google/uuid"
)

type Nasabah struct {
	ID        			uuid.UUID    `gorm:"primary_key;not_null" json:"id"`
	Nama 				string 		`json:"nama"`
	NoKtp 				string 		`json:"no_ktp"`
	TempatLahir 		string 		`json:"tempat_lahir"`
	TanggalLahir 		time.Time  	`json:"tanggal_lahir" time_format:"2006-01-02"`
	StatusKawin 		string 		`json:"status_kawin"`
	Alamat 				string 		`json:"alamat"`
	Pekerjaan 			string 		`json:"pekerjaan"`
	NamaPerusahaan 		string 		`json:"nama_perusahaan"`
	AlamatPerusahaan 	string 		`json:"alamat_perusahaan"`
	NoHp 				string 		`json:"no_hp"`
	Email 				string 		`json:"email" binding:"email"`
	NamaIbuKandung 		string 		`json:"nama_ibu_kandung"`
	JenisKelamin 		string 		`json:"jenis_kelamin"`
	
	CreatedAt 			time.Time 	`json:"created_at" default:"CURRENT_TIMESTAMP"`
	UpdatedAt 			time.Time 	`json:"updated_at"`
}