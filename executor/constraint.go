package executor

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	InvalidAffectValueErr = fmt.Errorf("db.Affect() only accept int type or string like 1+")
	RowsAffectedCheckErr  = fmt.Errorf("check affected rows error")
)

func newAffectingConstraint(v any) (a affectingConstraint, err error) {

	switch r := v.(type) {
	case int:
		return affectingConstraint{rows: r}, err
	case string:
		reg := regexp.MustCompile(`^([0-9]+)(\+)?$`)
		if !reg.MatchString(r) {
			err = fmt.Errorf("%w; got: %s", InvalidAffectValueErr, r)
			return
		}
		items := reg.FindStringSubmatch(r)
		var rows int64
		rows, err = strconv.ParseInt(items[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("db.Affect() parse value: %s to int error: %w", r, err)
			return
		}
		a.rows = int(rows)
		a.sign = items[2]
		return
	default:
		err = InvalidAffectValueErr
		return
	}
}

type affectingConstraint struct {
	rows int
	sign string
}

func (a affectingConstraint) Check(rows int) error {
	if a.sign != "" {
		if rows < a.rows {
			return fmt.Errorf("%w: expect affected rows >= %d, got %d", RowsAffectedCheckErr, a.rows, rows)
		}
	} else {
		if rows != a.rows {
			return fmt.Errorf("%w: expect affected rows = %d, got %d", RowsAffectedCheckErr, a.rows, rows)
		}
	}
	return nil
}
