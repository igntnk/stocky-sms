package repository

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"strconv"
)

func NumericToFloat64(n pgtype.Numeric) (float64, error) {
	val, err := n.Value()
	if err != nil {
		return 0, err
	}
	switch v := val.(type) {
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("неожиданный тип данных: %T", v)
	}
}

func Float64ToNumericWithPrecision(f float64) (pgtype.Numeric, error) {
	s := strconv.FormatFloat(f, 'g', -1, 64)

	n := pgtype.Numeric{}
	err := n.Scan(s)
	if err != nil {
		return pgtype.Numeric{}, fmt.Errorf("scan error: %w", err)
	}
	return n, nil
}
