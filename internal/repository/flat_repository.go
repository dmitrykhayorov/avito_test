package repository

import (
	"avito/internal/models"
	"database/sql"
	"errors"
	"fmt"
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

func (r *FlatRepository) GetFlatStatus(flatId int) (models.Status, error) {
	query := sq.Select("status").From("flat").
		Where(sq.Eq{"id": flatId}).PlaceholderFormat(sq.Dollar)

	var status models.Status
	row := query.RunWith(r.db).QueryRow()
	err := row.Scan(&status)
	if err != nil {
		return "", nil
	}
	return status, nil
}

func (r *FlatRepository) Update(flatId int, status models.Status) (models.Flat, error) {
	putOnModeration := sq.Update("flat").Set("status", models.OnModeration).
		Where(sq.Eq{"id": flatId}).PlaceholderFormat(sq.Dollar)

	_, err := putOnModeration.RunWith(r.db).Query()
	if err != nil {
		errorString := fmt.Sprintf("unable to put flat_id: %d on moderation", flatId)
		return models.Flat{}, errors.New(errorString)
	}

	updateStatusQuery := sq.Update("flat").Set("status", status).
		Where(sq.Eq{"id": flatId}).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)

	row := updateStatusQuery.RunWith(r.db).QueryRow()

	var updatedFlat models.Flat

	err = row.Scan(&updatedFlat.Number, &updatedFlat.HouseId, &updatedFlat.Price, &updatedFlat.Rooms,
		&updatedFlat.Status, &updatedFlat.CreatedAt)

	if err != nil {
		fmt.Println(err)
		return models.Flat{}, errors.New("unable to retrieve data from updated flat")
	}

	return updatedFlat, nil
}
