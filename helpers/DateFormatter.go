package helpers

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func StringToDate(value string) (time.Time, error) {

	parts := strings.Split(value, "-")
	if len(parts) != 3 {
		return time.Time{}, errors.New("bukan format tanggal")
	}

	day := parts[0]
	month := parts[1]
	year := parts[2]

	dateStr := fmt.Sprintf("%s-%s-%sT00:00:00Z", year, month, day)

	date, err := time.Parse(time.RFC3339, dateStr)

	if err != nil {
		fmt.Println("Gagal parsing tanggal:", err)
		return time.Time{}, err
	}

	return date, nil
}

func DateToString(value time.Time) string {
	year := value.Year()
	month := int(value.Month())
	day := value.Day()

	formattedDate := fmt.Sprintf("%02d-%02d-%d", day, month, year)

	return formattedDate
}
