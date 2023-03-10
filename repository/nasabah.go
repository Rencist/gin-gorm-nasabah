package repository

import (
	"context"
	"gin-gorm-nasabah/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NasabahRepository interface {
	GetAllNasabah(ctx context.Context) ([]entity.Nasabah, error)
	GetNasabahByID(ctx context.Context, NasabahID string) (entity.Nasabah, error)
	CreateNasabah(ctx context.Context, nasabah entity.Nasabah) (entity.Nasabah, error)
	UpdateNasabah(ctx context.Context, nasabah entity.Nasabah) (entity.Nasabah, error)
	DeleteNasabah(ctx context.Context, NasabahID string) (entity.Nasabah, error)
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
	tx := db.connection.Preload("Rekenings").Find(&listNasabah)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listNasabah, nil
}

func (db *nasabahConnection) GetNasabahByID(ctx context.Context, nasabahID string) (entity.Nasabah, error) {
	var nasabah entity.Nasabah
	nasabahUUID, err := uuid.Parse(nasabahID)
	if err != nil {
		return entity.Nasabah{}, err
	}
	nc := db.connection.Where("id = ?", nasabahUUID).Preload("Rekenings").Find(&nasabah)
	if nc.Error != nil {
		return entity.Nasabah{}, nc.Error
	}
	return nasabah, nil
}

func (db *nasabahConnection) CreateNasabah(ctx context.Context, nasabah entity.Nasabah) (entity.Nasabah, error) {
	nasabah.ID = uuid.New()
	nc := db.connection.Create(&nasabah)
	if nc.Error != nil {
		return entity.Nasabah{}, nc.Error
	}
	return nasabah, nil
}

func (db *nasabahConnection) UpdateNasabah(ctx context.Context, nasabah entity.Nasabah) (entity.Nasabah, error) {
	var detailNasabah entity.Nasabah
	nc := db.connection.Updates(&nasabah)
	if nc.Error != nil {
		return entity.Nasabah{}, nc.Error
	}
	res := db.connection.Where("id = ?", nasabah.ID).Preload("Rekenings").Find(&detailNasabah)
	if res.Error != nil {
		return entity.Nasabah{}, nc.Error
	}
	return detailNasabah, nil
} 

func (db *nasabahConnection) DeleteNasabah(ctx context.Context, nasabahID string) (entity.Nasabah, error) {
	nasabahUUID, err := uuid.Parse(nasabahID)
	if err != nil {
		return entity.Nasabah{}, err
	}
	nc := db.connection.Delete(&entity.Nasabah{}, &nasabahUUID)
	if nc.Error != nil {
		return entity.Nasabah{}, nc.Error
	}
	return entity.Nasabah{}, nil
}