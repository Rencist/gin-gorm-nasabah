package service

import (
	"context"
	"gin-gorm-nasabah/dto"
	"gin-gorm-nasabah/entity"
	"gin-gorm-nasabah/repository"

	"github.com/google/uuid"
)

type RekeningService interface {
	CreateRekening(ctx context.Context, rekeningDTO dto.RekeningCreateDto) (entity.Rekening, error)
	DeleteRekening(ctx context.Context, rekeningID string) (entity.Rekening, error)
}

type rekeningService struct {
	rekeningRepository repository.RekeningRepository
}

func NewRekeningService(rr repository.RekeningRepository) RekeningService {
	return &rekeningService{
		rekeningRepository: rr,
	}
}

func(rs *rekeningService) CreateRekening(ctx context.Context, rekeningDTO dto.RekeningCreateDto) (entity.Rekening, error) {
	nasabahUUID, _ := uuid.Parse(rekeningDTO.NasabahID)
	rekening := entity.Rekening{
		ID: rekeningDTO.ID,
		JenisAkun: rekeningDTO.JenisAkun,
		MataUang: rekeningDTO.MataUang,
		NasabahID: nasabahUUID,
	}
	return rs.rekeningRepository.CreateRekening(ctx, rekening)
}

func(rs *rekeningService) DeleteRekening(ctx context.Context, rekeningID string) (entity.Rekening, error) {
	return rs.rekeningRepository.DeleteRekening(ctx, rekeningID)
}