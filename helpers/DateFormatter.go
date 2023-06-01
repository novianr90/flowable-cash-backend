package helpers

import "time"

var dateLayout = "2006-01-02"

func StringToDate(value string) (time.Time, error) {

	date, err := time.Parse(dateLayout, value)

	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func DateToString(value time.Time) string {
	date := value.Format(dateLayout)

	return date
}
