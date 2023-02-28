package entity

import (
	"time"
)

type Rekening struct {
	ID        	uint64    	`gorm:"primary_key;auto_increment;not_null" json:"id"`
	JenisAkun 	string 		`json:"jenis_akun"`
	MataUang    string 		`json:"mata_uang"`
	CreatedAt 	time.Time 	`json:"created_at" default:"CURRENT_TIMESTAMP"`
	UpdatedAt 	time.Time 	`json:"updated_at"`

	NasabahID   uint64 		`gorm:"foreignKey" json:"nasabah_id"`
	Nasabah     *Nasabah  	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"nasabah,omitempty"`
}