package helpers

import "time"

func DateFormatter(value string) (time.Time, error) {
	dateLayout := "02/01/2006"

	date, err := time.Parse(dateLayout, value)

	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
