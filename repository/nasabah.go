package repository

import (
	"context"
	"encoding/json"
	"gin-gorm-nasabah/entity"

	"gorm.io/gorm"
)

type NasabahRepository interface {
	GetAllNasabah(ctx context.Context) ([]entity.Nasabah, error)
	CreateNasabah(ctx context.Context, nasabah entity.Nasabah) (entity.Nasabah, error)
}

type nasabahConnection struct {
	connection *gorm.DB
}

func NewNasabahRepository(db *gorm.DB) NasabahRepository {
	return &nasabahConnection{
		connection: db,
	}
}

func (db *nasabahConnection) GetAllNasabah(ctx context.Context) ([]entity.Nasabah, error) {
	var listNasabah []entity.Nasabah
	tx := db.connection.Debug().Find(&listNasabah)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listNasabah, nil
}

func (db *nasabahConnection) CreateNasabah(ctx context.Context, nasabah entity.Nasabah) (entity.Nasabah, error) {
	json.Marshal(nasabah.TanggalLahir)
	nc := db.connection.Create(&nasabah)
	if nc.Error != nil {
		return entity.Nasabah{}, nc.Error
	}
	return nasabah, nil
}