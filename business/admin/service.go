package admin

import (
	"api-jasa-pengiriman/business/ongkir"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindAdmins() (admins []Admin, err error)
	FindAdminByID(id int) (*Admin, error)
	InsertAdmin(admin *Admin) (*Admin, error)
	RemoveAdmin(id int) error
	RenewAdmin(id int, admin *Admin) (*Admin, error)
	FindAllCity() (city []ongkir.Kota, err error)
	InsertCity(kota *Kota) (*Kota, error)
	RenewCity(id int, datakota *Kota) (*Kota, error)
	RemoveCity(id int) (*Kota, error)
	CreateToken(admins *Admin) (string, error)
	FindCityByName(data *GetCityById) (*ongkir.Kota, error)
}

type Service interface {
	GetAdmins() (Admins []Admin, err error)
	GetAdminByID(id int) (*Admin, error)
	CreateAdmin(admin *Admin) (*Admin, error)
	DeleteAdmin(id int) error
	UpdateAdmin(id int, admin *Admin) (*Admin, error)
	GetAllCity() (city []ongkir.Kota, err error)
	GetToken(admins *Admin) (string, error)
	CreateCity(kota *Kota) (*Kota, error)
	UpdateCity(id int, datakota *Kota) (*Kota, error)
	DeleteCity(id int) (*Kota, error)
	GetCityByName(data *GetCityById) (*ongkir.Kota, error)
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

func (s *service) GetAdmins() (admins []Admin, err error) {
	admins, err = s.repository.FindAdmins()
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (s *service) GetAdminByID(id int) (*Admin, error) {
	return s.repository.FindAdminByID(id)
}

func (s *service) CreateAdmin(admin *Admin) (*Admin, error) {
	err := s.validate.Struct(admin)
	if err != nil {
		return nil, err
	}
	admin, err = s.repository.InsertAdmin(admin)
	return admin, err
}

func (s *service) DeleteAdmin(id int) error {
	return s.repository.RemoveAdmin(id)
}

func (s *service) UpdateAdmin(id int, admin *Admin) (*Admin, error) {
	return s.repository.RenewAdmin(id, admin)
}

func (s *service) GetAllCity() (city []ongkir.Kota, err error) {
	return s.repository.FindAllCity()
}

func (s *service) GetToken(admins *Admin) (string, error) {
	tokens, err := s.repository.CreateToken(admins)
	return tokens, err
}

func (s *service) CreateCity(kota *Kota) (*Kota, error) {
	err := s.validate.Struct(kota)
	if err != nil {
		return nil, err
	}
	return s.repository.InsertCity(kota)
}

func (s *service) UpdateCity(id int, datakota *Kota) (*Kota, error) {
	return s.repository.RenewCity(id, datakota)
}

func (s *service) DeleteCity(id int) (*Kota, error) {
	return s.repository.RemoveCity(id)
}

func (s *service) GetCityByName(data *GetCityById) (*ongkir.Kota, error) {
	return s.repository.FindCityByName(data)
}
