package gobatis

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Scanner interface {
	Scan(rows *sql.Rows, ct *sql.ColumnType, value reflect.Value) error
}

type ScannerFactory func() Scanner

type queryScanner struct {
	rows     *sql.Rows
	first    bool
	mapping  bool
	selected map[string]int
	values   []reflect.Value
	scanTag  string
	scanner  Scanner
}

func (p *queryScanner) Rows() *sql.Rows {
	return p.rows
}

func (p *queryScanner) setSelected(rt int, params []*param, values []reflect.Value) error {
	
	p.values = values
	p.mapping = len(params) == 0
	
	if rt != result_result {
		return nil
	}
	
	var el int
	if p.mapping {
		el = 1
	} else {
		el = len(params)
	}
	
	if el != len(values) {
		return fmt.Errorf("expected to receive %d result filed(s), got %d (except error)", el, len(values))
	}
	
	var err error
	if p.mapping {
		elem := reflectValueElem(p.values[0])
		if elem.Kind() == reflect.Struct {
			for i := 0; i < elem.NumField(); i++ {
				field := trimScanComma(elem.Type().Field(i).Tag.Get(p.scanTag))
				if field != "" {
					p.values = append(p.values, elem.Field(i))
					err = p.addSelected(len(p.values)-1, field)
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	}
	p.first = true
	for i, v := range params {
		err = p.addSelected(i, v.name)
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

func (p *queryScanner) scan() (err error) {
	cts, err := p.rows.ColumnTypes()
	if err != nil {
		return
	}
	i := -1
	
	if p.mapping {
		
		for p.rows.Next() {
			i++
			if p.isSelected(cts[i].Name()) {
				err = p.scanner.Scan(p.rows, cts[i], p.values[p.selected[cts[i].Name()]])
				if err != nil {
					return
				}
				if p.first {
					break
				}
			}
			
		}
	} else {
		for p.rows.Next() {
			i++
			err = p.scanner.Scan(p.rows, cts[i], p.values[i])
			if err != nil {
				return
			}
			if p.first {
				break
			}
		}
	}
	if len(p.values) > 0 && i == -1 {
		err = sql.ErrNoRows
		return
	}
	return
}

type execScanner struct {
	affected int64
	values   []reflect.Value
}

func (p *execScanner) scan() error {
	if len(p.values) != 1 {
		return fmt.Errorf("scan rows: illegal values length %d", len(p.values))
	}
	// TODO handle pointer
	p.values[0].Elem().SetInt(p.affected)
	//switch p.values[0].Elem().Kind() {
	//case reflect.Int:
	//	r, e := cast.ToIntE(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetInt(int64(r))
	//case reflect.Int8:
	//	r, e := cast.ToInt8E(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetInt(int64(r))
	//case reflect.Int16:
	//	r, e := cast.ToInt16E(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetInt(int64(r))
	//case reflect.Int32:
	//	r, e := cast.ToInt32E(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetInt(int64(r))
	//case reflect.Int64:
	//	p.values[0].Elem().SetInt(p.affected)
	//case reflect.Uint:
	//	r, e := cast.ToUintE(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetUint(uint64(r))
	//case reflect.Uint8:
	//	r, e := cast.ToUint8E(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetUint(uint64(r))
	//case reflect.Uint16:
	//	r, e := cast.ToUint16E(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetUint(uint64(r))
	//case reflect.Uint32:
	//	r, e := cast.ToUint32E(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetUint(uint64(r))
	//case reflect.Uint64:
	//	r, e := cast.ToUint64E(p.affected)
	//	if e != nil {
	//		return e
	//	}
	//	p.values[0].Elem().SetUint(r)
	//}
	return nil
}
