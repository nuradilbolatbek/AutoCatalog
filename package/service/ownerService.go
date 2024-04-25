package service

import (
	"autokatolog"
	"autokatolog/package/repository"
)

type ownerService struct {
	repo repository.OwnerRepository
}

func NewOwnerService(repo repository.OwnerRepository) OwnerService {
	return &ownerService{repo: repo}
}

func (s *ownerService) CreateOwner(owner autokatolog.People) (int, error) {
	return s.repo.CreateOwner(owner)
}
