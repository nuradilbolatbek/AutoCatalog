package repository

import (
	"autokatolog"
	"github.com/jmoiron/sqlx"
)

type CarRepository interface {
	Create(regNum autokatolog.Car) error
	GetAllCars(filter autokatolog.Car, page, pageSize int) ([]autokatolog.Car, error)
	GetByRegNum(regNum string) (autokatolog.Car, error)
	Update(car autokatolog.Car) error
	Delete(regNum string) error
}

type OwnerRepository interface {
	CreateOwner(owner autokatolog.People) (int, error)
}

type Repository struct {
	CarRepository
	OwnerRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CarRepository:   NewCarRepo(db),
		OwnerRepository: NewOwnerRepo(db),
	}

}
