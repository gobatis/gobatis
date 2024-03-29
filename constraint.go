package batis

import (
	"fmt"
	"regexp"
	"strconv"
)

func newAffectConstraint(v any) (*affectConstraint, error) {

	switch r := v.(type) {
	case int:
		return &affectConstraint{rows: int64(r)}, nil
	case int64:
		return &affectConstraint{rows: int64(r)}, nil
	case string:
		reg := regexp.MustCompile(`^([0-9]+)(\+)?$`)
		if !reg.MatchString(r) {
			return nil, fmt.Errorf("%w; got: %s", ErrInvalidAffectValue, r)
		}
		items := reg.FindStringSubmatch(r)
		var rows int64
		rows, err := strconv.ParseInt(items[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("db.Affect() parse value: %s to int error: %w", r, err)
		}
		return &affectConstraint{rows: int64(rows), sign: items[2]}, nil
	default:
		return nil, ErrInvalidAffectValue
	}
}

type affectConstraint struct {
	rows int64
	sign string
}

func (a affectConstraint) Check(rows int64) error {
	if a.sign != "" {
		if rows < a.rows {
			return fmt.Errorf("%w, expect affected rows >= %d, got %d", ErrAffectConstrict, a.rows, rows)
		}
	} else {
		if rows != a.rows {
			return fmt.Errorf("%w, expect affected rows = %d, got %d", ErrAffectConstrict, a.rows, rows)
		}
	}
	return nil
}
