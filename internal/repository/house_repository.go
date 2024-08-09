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

func (r *HouseRepository) Create(house models.House) (models.House, error) {
	//h := sq.Insert("house").Columns("address", "year", "developer").
	//	Values(house.Address, house.Year, house.Developer)

	// TODO: add transactions
	query := `INSERT INTO house (address,year,developer) VALUES ($1, $2, $3) RETURNING *`

	err := r.db.QueryRow(query, house.Address, house.Year, house.Developer).
		Scan(&house.Id, &house.Address, &house.Year, &house.Developer, &house.CreatedAt, &house.UpdatedAt)

	if err != nil {
		// TODO: change it for empty struct
		return house, err
	}

	return house, nil
}

func (r *HouseRepository) GetFlatsByHouseID(userRole models.UserRole, houseId uint32) ([]models.Flat, error) {
	// TODO: add transactions
	query := sq.Select("*").From("flat").Where(sq.Eq{"house_id": houseId})
	if userRole != models.Moderator {
		query = query.Where(sq.Eq{"status": models.Approved})
	}
	query = query.RunWith(r.db).PlaceholderFormat(sq.Dollar)

	q, _, _ := query.ToSql()
	fmt.Println(q)
	rows, err := query.Query()

	if err != nil {
		return nil, err
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

	return flats, nil
}
