package dto

import "github.com/google/uuid"

type RekeningCreateDto struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id" form:"id"`
	JenisAkun string 	`json:"jenis_akun" binding:"required"`
	MataUang  string 	`json:"mata_uang" binding:"required"`

	NasabahID uuid.UUID `gorm:"foreignKey" json:"nasabah_id" binding:"required"`
}