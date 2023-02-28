package entity

import (
	"time"

	"github.com/google/uuid"
)

type Rekening struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	JenisAkun 	string 		`json:"jenis_akun"`
	MataUang    string 		`json:"mata_uang"`
	CreatedAt 	time.Time 	`json:"created_at" default:"CURRENT_TIMESTAMP"`
	UpdatedAt 	time.Time 	`json:"updated_at"`

	NasabahID   uuid.UUID 		`gorm:"foreignKey" json:"nasabah_id"`
	Nasabah     *Nasabah  	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"nasabah,omitempty"`
}