package repository

import (
	"avito/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type FlatRepositoryInterface interface {
	Create(flat models.Flat) (models.Flat, error)
	GetFlatStatus(flatId int) (models.Status, error)
	Update(flatId int, houseId int, status models.Status) (models.Flat, error)
}

type FlatRepository struct {
	db *sql.DB
}

func NewFlatRepository(db *sql.DB) *FlatRepository {
	return &FlatRepository{
		db: db,
	}
}

func (r *FlatRepository) Create(flat models.Flat) (models.Flat, error) {
	const op = "flatRepository.Create"

	fail := func(err error) (models.Flat, error) {
		return models.Flat{}, fmt.Errorf("op: %v, err: %v", op, err)
	}

	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer tx.Rollback()

	columns := []string{"house_id", "price"}
	values := []interface{}{flat.HouseId, flat.Price}

	if flat.Rooms != nil {
		columns = append(columns, "rooms")
		values = append(values, flat.Rooms)
	}

	query := sq.Insert("flat").Columns(columns...).
		Values(values...).
		Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)

	rows, err := query.RunWith(r.db).Query()

	if err != nil {
		return fail(err)
	}

	defer rows.Close()

	rows.Next()
	if err := rows.Scan(&flat.Id, &flat.HouseId, &flat.Price,
		&flat.Rooms, &flat.Status, &flat.CreatedAt); err != nil {
		return models.Flat{}, err
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return flat, nil
}

func (r *FlatRepository) GetFlatStatus(flatId int) (models.Status, error) {
	const op = "flatRepository.GetFlatStatus"

	fail := func(err error) (models.Status, error) {
		return "", fmt.Errorf("op: %v, err: %v", op, err)
	}

	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer tx.Rollback()

	query := sq.Select("status").From("flat").
		Where(sq.Eq{"id": flatId}).PlaceholderFormat(sq.Dollar)

	var status models.Status
	row := query.RunWith(r.db).QueryRow()
	err = row.Scan(&status)

	if err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return status, nil
}

func (r *FlatRepository) Update(flatId int, houseId int, status models.Status) (models.Flat, error) {
	const op = "flatRepository.Update"

	fail := func(err error) (models.Flat, error) {
		return models.Flat{}, fmt.Errorf("op: %v, err: %v", op, err)
	}

	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	putOnModeration := sq.Update("flat").Set("status", models.StatusOnModeration).
		Where(sq.Eq{"id": flatId, "house_id": houseId}).PlaceholderFormat(sq.Dollar)

	_, err = putOnModeration.RunWith(r.db).Query()
	if err != nil {
		errorString := fmt.Sprintf("unable to put flat_id: %d on moderation", flatId)
		return fail(errors.New(errorString))
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	tx, err = r.db.BeginTx(ctx, nil)

	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	updateStatusQuery := sq.Update("flat").Set("status", status).
		Where(sq.Eq{"id": flatId, "house_id": houseId}).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)

	row := updateStatusQuery.RunWith(r.db).QueryRow()

	var updatedFlat models.Flat

	err = row.Scan(&updatedFlat.Id, &updatedFlat.HouseId, &updatedFlat.Price, &updatedFlat.Rooms,
		&updatedFlat.Status, &updatedFlat.CreatedAt)

	if err != nil {
		return fail(errors.New("unable to retrieve data from updated flat:" + err.Error()))
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return updatedFlat, nil
}
