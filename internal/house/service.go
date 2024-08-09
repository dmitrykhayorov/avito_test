package house

import (
	"avito/internal/models"
	"errors"
)

func validateHouseData(house models.House) error {
	if house.Address == "" {
		return errors.New("empty address")
	}
	if house.Year < 0 {
		return errors.New("year is less than 0")
	}
	return nil
}
