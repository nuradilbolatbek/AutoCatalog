package repository

import (
	"autokatolog"
	"github.com/jmoiron/sqlx"
	"log"
)

type OwnerRepo struct {
	db *sqlx.DB
}

func NewOwnerRepo(db *sqlx.DB) *OwnerRepo {
	return &OwnerRepo{db: db}
}

// CreateOwner inserts a new owner into the database and returns the new owner's ID.
func (r *OwnerRepo) CreateOwner(owner autokatolog.People) (int, error) {
	var ownerId int
	query := `INSERT INTO owners (name, surname, patronymic) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, owner.Name, owner.Surname, owner.Patronymic).Scan(&ownerId)
	if err != nil {
		log.Printf("Failed to create owner: %s", err)
		return 0, err
	}
	return ownerId, nil
}
