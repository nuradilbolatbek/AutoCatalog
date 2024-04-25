package repository

import (
	"autokatolog"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type CarRepo struct {
	db *sqlx.DB
}

func NewCarRepo(db *sqlx.DB) *CarRepo {
	return &CarRepo{db: db}
}

func (r *CarRepo) Create(car autokatolog.Car) error {
	query := `INSERT INTO cars (reg_num, mark, model, year, owner_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, car.RegNums, car.Mark, car.Model, car.Year, car.Owner.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CarRepo) GetAllCars(filter autokatolog.Car, page, pageSize int) ([]autokatolog.Car, error) {
	var cars []autokatolog.Car
	args := []interface{}{}
	whereClauses := []string{"1 = 1"}

	if filter.RegNums != "" {
		whereClauses = append(whereClauses, "c.reg_num = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.RegNums)
	}
	if filter.Mark != "" {
		whereClauses = append(whereClauses, "c.mark = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Mark)
	}
	if filter.Model != "" {
		whereClauses = append(whereClauses, "c.model = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Model)
	}
	if filter.Year != 0 {
		whereClauses = append(whereClauses, "c.year = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Year)
	}
	if filter.Owner.ID != 0 {
		whereClauses = append(whereClauses, "o.id = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Owner.ID)
	}

	query := fmt.Sprintf(`
    SELECT c.id, c.reg_num, c.mark, c.model, c.year, o.id, o.name, o.surname, o.patronymic
    FROM cars c
    JOIN owners o ON c.owner_id = o.id
    WHERE %s
    LIMIT $%d OFFSET $%d`, strings.Join(whereClauses, " AND "), len(args)+1, len(args)+2)

	args = append(args, pageSize, (page-1)*pageSize)

	fmt.Printf("Executing query: %s\nWith args: %+v\n", query, args)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query execution error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var car autokatolog.Car
		err := rows.Scan(&car.ID, &car.RegNums, &car.Mark, &car.Model, &car.Year, &car.Owner.ID, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (r *CarRepo) GetByRegNum(regNum string) (autokatolog.Car, error) {
	var car autokatolog.Car
	query := `
SELECT c.id, c.reg_num, c.mark, c.model, c.year, 
       o.id as "owner.id", o.name as "owner.name", 
       o.surname as "owner.surname", o.patronymic as "owner.patronymic"
FROM cars c
JOIN owners o ON c.owner_id = o.id
WHERE c.reg_num = $1`
	err := r.db.Get(&car, query, regNum)
	if err != nil {
		return autokatolog.Car{}, err
	}
	return car, nil
}

func (r *CarRepo) Update(car autokatolog.Car) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	fmt.Println(car)

	updateOwnerQuery := `UPDATE owners SET name = $1, surname = $2, patronymic = $3 WHERE id = $4`
	if _, err := tx.Exec(updateOwnerQuery, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic, car.Owner.ID); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(car)

	updateCarQuery := `UPDATE cars SET mark = $1, model = $2, year = $3 WHERE reg_num = $4`
	if _, err := tx.Exec(updateCarQuery, car.Mark, car.Model, car.Year, car.RegNums); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *CarRepo) Delete(regNum string) error {
	query := `DELETE FROM cars WHERE reg_num = $1`
	_, err := r.db.Exec(query, regNum)
	return err
}
