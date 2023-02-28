package dto

type RekeningCreateDto struct {
	JenisAkun string    `json:"jenis_akun" binding:"required"`
	MataUang  string    `json:"mata_uang" binding:"required"`

	NasabahID uint64 `gorm:"foreignKey" json:"nasabah_id" binding:"required"`
}