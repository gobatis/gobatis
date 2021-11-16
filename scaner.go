package gobatis

import (
	"database/sql"
	"fmt"
	"github.com/gobatis/gobatis/cast"
	"reflect"
	"strings"
)

type Scanner interface {
	Scan(rows *sql.Rows, ct *sql.ColumnType, value reflect.Value) error
}

type ScannerFactory func() Scanner

type queryScanner struct {
	rows     *sql.Rows
	first    bool
	reflect  bool
	all      bool
	selected map[string]int
	values   []reflect.Value
	tag      string
	scanner  Scanner
}

func (p *queryScanner) Rows() *sql.Rows {
	return p.rows
}

func (p *queryScanner) setSelected(rt int, params []*param, values []reflect.Value) error {
	
	p.values = values
	p.reflect = len(params) == 0
	
	if rt != result_result {
		return nil
	}
	
	var el int
	if p.reflect {
		el = 1
	} else {
		el = len(params)
	}
	
	if el != len(values) {
		return fmt.Errorf("expected to receive %d result filed(s), got %d (except error)", el, len(values))
	}
	
	if p.reflect {
		return nil
	}
	
	p.first = true
	for i, v := range params {
		err := p.addSelected(i, v.name)
		if err != nil {
			return err
		}
		if v.slice {
			p.first = false
		}
	}
	
	return nil
}

func (p *queryScanner) addSelected(index int, name string) error {
	if p.selected == nil {
		p.selected = map[string]int{}
	}
	if _, ok := p.selected[name]; ok {
		return fmt.Errorf("duplicated result filed '%s'", name)
	}
	p.selected[name] = index
	return nil
}

func (p *queryScanner) isSelected(field string) bool {
	if p.selected == nil {
		return false
	}
	_, ok := p.selected[field]
	return ok
}

type Assigner interface {
	AssignTo(interface{}) error
}

func (p *queryScanner) scan() (err error) {
	cts, err := p.rows.ColumnTypes()
	if err != nil {
		return
	}
	i := 0
	if p.reflect {
		for p.rows.Next() {
			if p.isSelected(cts[i].Name()) {
				err = p.scanner.Scan(p.rows, cts[i], p.values[p.selected[cts[i].Name()]])
				if err != nil {
					return
				}
				if p.first {
					break
				}
			}
			i++
		}
	} else {
		for p.rows.Next() {
			err = p.scanner.Scan(p.rows, cts[i], p.values[i])
			if err != nil {
				return
			}
			if p.first {
				break
			}
			i++
		}
	}
	if len(p.values) > 0 && i == 0 {
		err = sql.ErrNoRows
		return
	}
	return
}

func (p *queryScanner) trimComma(field string) string {
	// TODO 也许可以更加优化
	if strings.Contains(field, ",") {
		return strings.TrimSpace(strings.Split(field, ",")[0])
	}
	return field
}

type execScanner struct {
	affected int64
	values   []reflect.Value
}

func (p *execScanner) scan() error {
	if len(p.values) != 0 {
		return fmt.Errorf("scan rows: illegal values length %d", len(p.values))
	}
	switch p.values[0].Elem().Kind() {
	case reflect.Int:
		r, e := cast.ToIntE(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetInt(int64(r))
	case reflect.Int8:
		r, e := cast.ToInt8E(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetInt(int64(r))
	case reflect.Int16:
		r, e := cast.ToInt16E(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetInt(int64(r))
	case reflect.Int32:
		r, e := cast.ToInt32E(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetInt(int64(r))
	case reflect.Int64:
		p.values[0].Elem().SetInt(p.affected)
	case reflect.Uint:
		r, e := cast.ToUintE(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetUint(uint64(r))
	case reflect.Uint8:
		r, e := cast.ToUint8E(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetUint(uint64(r))
	case reflect.Uint16:
		r, e := cast.ToUint16E(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetUint(uint64(r))
	case reflect.Uint32:
		r, e := cast.ToUint32E(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetUint(uint64(r))
	case reflect.Uint64:
		r, e := cast.ToUint64E(p.affected)
		if e != nil {
			return e
		}
		p.values[0].Elem().SetUint(r)
	}
	
	return nil
}
