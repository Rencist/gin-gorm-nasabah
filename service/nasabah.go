package service

import (
	"context"
	"gin-gorm-nasabah/entity"
	"gin-gorm-nasabah/repository"
)

type NasabahService interface {
	GetAllNasabah(ctx context.Context) ([]entity.Nasabah, error)
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