package service

import (
	"context"
	"gin-gorm-nasabah/dto"
	"gin-gorm-nasabah/entity"
	"gin-gorm-nasabah/repository"

	"github.com/mashingan/smapping"
)

type NasabahService interface {
	GetAllNasabah(ctx context.Context) ([]entity.Nasabah, error)
	CreateNasabah(ctx context.Context, nasabahDTO dto.NasabahCreateDto) (entity.Nasabah, error)
}

type nasabahService struct {
	nasabahRepository repository.NasabahRepository
}

func NewNasabahService(nr repository.NasabahRepository) NasabahService {
	return &nasabahService{
		nasabahRepository: nr,
	}
}

func(ns *nasabahService) GetAllNasabah(ctx context.Context) ([]entity.Nasabah, error) {
	return ns.nasabahRepository.GetAllNasabah(ctx)
}

func(ns *nasabahService) CreateNasabah(ctx context.Context, nasabahDTO dto.NasabahCreateDto) (entity.Nasabah, error) {
	nasabah := entity.Nasabah{}
	err := smapping.FillStruct(&nasabah, smapping.MapFields(nasabahDTO))
	if err != nil {
		return nasabah, err
	}
	return ns.nasabahRepository.CreateNasabah(ctx, nasabah)
}