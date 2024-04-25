package service

import (
	"autokatolog"
	"autokatolog/package/repository"
)

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

func (s *carService) CreateCar(regNum autokatolog.Car) error {
	return s.repo.Create(regNum)
}

func (s *carService) GetAllCars(filter autokatolog.Car, page, pageSize int) ([]autokatolog.Car, error) {
	return s.repo.GetAllCars(filter, page, pageSize)
}

func (s *carService) GetCarByRegNum(regNum string) (autokatolog.Car, error) {
	return s.repo.GetByRegNum(regNum)
}

func (s *carService) UpdateCar(car autokatolog.Car) (autokatolog.Car, error) {
	err := s.repo.Update(car)
	if err != nil {
		return autokatolog.Car{}, err
	}
	return car, nil
}

func (s *carService) DeleteCar(regNum string) error {
	return s.repo.Delete(regNum)
}
