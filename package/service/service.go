package service

import (
	"autokatolog"
	"autokatolog/package/repository"
)

type CarService interface {
	CreateCar(regNum autokatolog.Car) error
	GetAllCars(filter autokatolog.Car, page, pageSize int) ([]autokatolog.Car, error)
	GetCarByRegNum(regNum string) (autokatolog.Car, error)
	UpdateCar(car autokatolog.Car) (autokatolog.Car, error)
	DeleteCar(regNum string) error
}

type OwnerService interface {
	CreateOwner(owner autokatolog.People) (int, error)
}

type Service struct {
	CarService
	OwnerService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		CarService:   NewCarService(repo.CarRepository),
		OwnerService: NewOwnerService(repo.OwnerRepository),
	}
}
