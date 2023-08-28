package validation

import (
	"errors"
	"time"
)

func ValidateDate(date string) error {
	_, err := time.Parse("2006-01-02", date)
	if err == nil {
		return nil
	}
	_, err = time.Parse("2006-01", date)
	return err
}

func ValidateSegment(name string) error {
	if len(name) > 32 {
		return errors.New("name length must be less than 33")
	}
	return nil
}

func ValidatePercent(percent float64) error {
	if percent > 1 || percent < 0 {
		return errors.New("percent must be from 0.0 to 1.0")
	}
	return nil
}
