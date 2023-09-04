package batis

import (
	"database/sql"
)

func newClosure(db *sql.DB, tx *tx, executors ...*executor) *closure {
	return &closure{
		db:        db,
		tx:        tx,
		executors: executors,
	}
}

type closure struct {
	executors []*executor
	tx        *tx
	db        *sql.DB
}

func (c closure) conn() conn {
	//if d.tx != nil {
	//	c = d.tx
	//	return
	//} else {
	//	var err error
	//	c, err = d.db.Conn(d.context())
	//	if err != nil {
	//		d.addError(err)
	//		return
	//	}
	//	return
	//}
	return nil
}

func (c closure) exec() (err error) {
	for _, e := range c.executors {
		//e.conn = c.conn
		err = e.execute()
		if err != nil {
			return
		}
	}
	return
}

type Scanner struct {
	Error  error
	rows   *sql.Rows
	result sql.Result
}
