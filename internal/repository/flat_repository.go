package repository

import (
	"avito/internal/models"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
)

type FlatRepository struct {
	db *sql.DB
}

func NewFlatRepository(db *sql.DB) *FlatRepository {
	return &FlatRepository{
		db: db,
	}
}

func (r *FlatRepository) Create(flat models.Flat) (models.Flat, error) {
	query := sq.Insert("flat").Columns("house_id", "price", "rooms", "status").
		Values(flat.HouseId, flat.Price, flat.Rooms, models.Created).
		Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)

	rows, err := query.RunWith(r.db).Query()

	if err != nil {
		return models.Flat{}, err
	}

	rows.Next()
	if err := rows.Scan(&flat.Number, &flat.HouseId, &flat.Price,
		&flat.Rooms, &flat.Status, &flat.CreatedAt); err != nil {
		return models.Flat{}, err
	}

	return flat, nil
}

func (r *FlatRepository) Update(flatId int) {

}
