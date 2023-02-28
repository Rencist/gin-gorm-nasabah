package entity

import (
	"time"
)

type Rekening struct {
	ID        	uint64    	`gorm:"primary_key;auto_increment;not_null" json:"id" binding:"required"`
	JenisAkun 	string 		`json:"jenis_akun" binding:"required"`
	MataUang    string 		`json:"mata_uang" binding:"required"`
	CreatedAt 	time.Time 	`json:"created_at" default:"CURRENT_TIMESTAMP" binding:"required"`
	UpdatedAt 	time.Time 	`json:"updated_at" binding:"required"`

	NasabahID   uint64 		`gorm:"foreignKey" json:"nasabah_id" binding:"required"`
	Nasabah     *Nasabah  	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"nasabah,omitempty"`
}