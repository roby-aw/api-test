package ongkir

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	Cost(Dataongkir *Ongkir) (*CostResponse, error)
	Resi(Dataresi *Resi) (*CekResiBinderByte, error)
}

type Service interface {
	GetCost(Dataongkir *Ongkir) (*CostResponse, error)
	GetResi(Dataresi *Resi) (*CekResiBinderByte, error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) GetCost(Dataongkir *Ongkir) (*CostResponse, error) {
	return s.repository.Cost(Dataongkir)
}

func (s *service) GetResi(Dataresi *Resi) (*CekResiBinderByte, error) {
	err := s.validate.Struct(Dataresi)
	resi, err := s.repository.Resi(Dataresi)
	if err != nil {
		return nil, err
	}
	if resi.Message == "Parameters `courier` and `awb` is required" {
		return nil, fmt.Errorf("input kurir atau resi yang dimasukkan salah")
	}
	return resi, nil
}
