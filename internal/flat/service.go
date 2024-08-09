package flat

import (
	"avito/internal/models"
	"errors"
)

func validateFlatData(flat models.Flat) error {
	if flat.Price == nil {
		return errors.New("missing price")
	}

	if *flat.Price < 0 {
		return errors.New("price is less than 0")
	}

	if flat.HouseId < 1 {
		return errors.New("invalid house_id")
	}

	if flat.Rooms != nil && *flat.Rooms < 1 {
		return errors.New("number of rooms is less that 1")
	}

	return nil
}
