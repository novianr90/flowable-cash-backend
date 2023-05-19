package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Balance struct {
	Debit  float64
	Credit float64
}

func (b *Balance) Scan(value interface{}) error {

	strValue, ok := value.(string)

	if !ok {
		return errors.New("invalid value type. expected string")
	}

	parts := strings.Split(strValue, ":")

	debit, err := strconv.ParseFloat(parts[0], 64)

	if err != nil {
		return errors.New("failed to parse debit value")
	}

	credit, err := strconv.ParseFloat(parts[1], 64)

	if err != nil {
		return errors.New("failed to parse credit value")
	}

	b.Debit = debit
	b.Credit = credit

	return nil
}

func (b *Balance) Value() (driver.Value, error) {
	value := fmt.Sprintf("%.2f:%.2f", b.Debit, b.Credit)

	return value, nil
}
