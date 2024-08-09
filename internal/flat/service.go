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

func validateStatus(status models.Status) error {
	validStatuses := map[models.Status]struct{}{
		models.Created:      {},
		models.Approved:     {},
		models.Declined:     {},
		models.OnModeration: {},
	}

	_, ok := validStatuses[status]
	if !ok {
		return errors.New("unsupported status")
	}

	if status == models.OnModeration {
		return errors.New("status on moderation will be applied automatically, specify other status")
	}

	return nil
}
