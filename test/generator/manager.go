package generator

import "sync"

type DataRow struct {
	Id     int
	Source string
	Item   interface{}
}

func NewDataManager() *DataManager {
	return &DataManager{
		rows: map[int]*DataRow{},
	}
}

type DataManager struct {
	id   int
	mu   sync.Mutex
	rows map[int]*DataRow
}

func (p *DataManager) NextID() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.id++
	return p.id
}

func (p *DataManager) AddRow(id int, source string, item interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.rows[id] = &DataRow{
		Id:     id,
		Source: source,
		Item:   item,
	}
}

func (p *DataManager) GetRow(id int) *DataRow {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.rows[id]
}
