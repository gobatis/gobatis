package generator

import (
	"fmt"
	"sync"
)

type SID struct {
	id int
	mu sync.Mutex
}

func (p *SID) NextId() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.id++
	return fmt.Sprintf("%d", p.id)
}

type DataRow struct {
	SId    string
	Source string
	Item   interface{}
}

func NewDataManager() *DataManager {
	return &DataManager{
		rows: map[string]*DataRow{},
	}
}

type DataManager struct {
	sid  int
	mu   sync.Mutex
	rows map[string]*DataRow
}

func (p *DataManager) NextId() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.sid++
	return fmt.Sprintf("%d", p.sid)
}

func (p *DataManager) AddRow(sid string, source string, item interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.rows[sid] = &DataRow{
		SId:    sid,
		Source: source,
		Item:   item,
	}
}

func (p *DataManager) GetRow(sid string) *DataRow {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.rows[sid]
}
