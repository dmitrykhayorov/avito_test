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

func validateUpdateRequestBody(body models.FlatUpdateRequestBody) error {
	validStatuses := map[models.Status]struct{}{
		models.StatusCreated:      {},
		models.StatusApproved:     {},
		models.StatusDeclined:     {},
		models.StatusOnModeration: {},
	}

	_, ok := validStatuses[body.Status]
	if !ok {
		return errors.New("unsupported status")
	}

	if body.Status == models.StatusOnModeration {
		return errors.New("status on moderation will be applied automatically, specify other status")
	}

	if body.FlatId < 1 {
		return errors.New("flat_id is less than 1")
	}

	if body.HouseId < 1 {
		return errors.New("house_id is less than 1")
	}

	return nil
}
