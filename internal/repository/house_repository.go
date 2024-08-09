package repository

import (
	"avito/internal/models"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type HouseRepository struct {
	db *sql.DB
}

func NewHouseRepository(db *sql.DB) *HouseRepository {
	return &HouseRepository{db: db}
}

func (r *HouseRepository) Create(house *models.House) (*models.House, error) {
	h := sq.Insert("house").Columns("address", "year", "developer").
		Values(house.Address, house.Year, house.Developer)

	sqll, _, _ := h.ToSql()
	fmt.Println(sqll, house.Address, house.Year, house.Developer)

	rows, err := r.db.Query(sqll, house.Address, house.Year, house.Developer)
	if err != nil {
		return house, err
	}

	if err != nil {
		// TODO: change it for empty struct
		return house, err
	}
	defer rows.Close()

	err = rows.Scan(&house.Id, &house.Address, &house.Year, &house.Developer, &house.CreatedAt, &house.UpdatedAt)

	if err != nil {
		return house, err
	}

	return house, nil
}

func (r *HouseRepository) GetFlatsByHouseID() []models.Flat {

	return nil
}
