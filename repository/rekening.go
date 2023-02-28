package repository

import (
	"context"
	"gin-gorm-nasabah/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RekeningRepository interface {
	CreateRekening(ctx context.Context, rekening entity.Rekening) (entity.Rekening, error)
}

type rekeningConnection struct {
	connection *gorm.DB
}

func NewRekeningRepository(db *gorm.DB) RekeningRepository {
	return &rekeningConnection{
		connection: db,
	}
}

func (db *rekeningConnection) CreateRekening(ctx context.Context, rekening entity.Rekening) (entity.Rekening, error) {
	rekening.ID = uuid.New()
	rc := db.connection.Create(&rekening)
	if rc.Error != nil {
		return entity.Rekening{}, rc.Error
	}
	return rekening, nil
}