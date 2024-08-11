package repository

import (
	"avito/internal/models"
	"context"
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

func (r *HouseRepository) Create(house models.House) (models.House, error) {
	const op = "HouseRepository.Create"

	fail := func(err error) (models.House, error) {
		return models.House{}, fmt.Errorf("op: %v, err: %v", op, err)
	}

	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	query := `INSERT INTO house (address,year,developer) VALUES ($1, $2, $3) RETURNING *`

	err = r.db.QueryRow(query, house.Address, house.Year, house.Developer).
		Scan(&house.Id, &house.Address, &house.Year, &house.Developer, &house.CreatedAt, &house.UpdatedAt)

	if err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return house, nil
}

func (r *HouseRepository) GetFlatsByHouseID(userRole models.UserRole, houseId uint32) ([]models.Flat, error) {
	const op = "HouseRepository.GetFlatsByHouseID"

	fail := func(err error) ([]models.Flat, error) {
		return nil, fmt.Errorf("op: %v, err: %v", op, err)
	}

	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	query := sq.Select("*").From("flat").Where(sq.Eq{"house_id": houseId})
	if userRole != models.Moderator {
		query = query.Where(sq.Eq{"status": models.Approved})
	}
	query = query.RunWith(r.db).PlaceholderFormat(sq.Dollar)

	rows, err := query.Query()

	if err != nil {
		return fail(err)
	}
	defer rows.Close()

	flats := make([]models.Flat, 0)

	for rows.Next() {
		var placeholder models.Flat
		if err := rows.Scan(&placeholder.Number, &placeholder.HouseId, &placeholder.Price,
			&placeholder.Rooms, &placeholder.Status, &placeholder.CreatedAt); err != nil {
			return nil, err
		}
		flats = append(flats, placeholder)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return flats, nil
}
