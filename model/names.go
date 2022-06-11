package model

import "sync/atomic"

// TODO: refactor to struct with mysql dependency

var m = map[uint32]string{}
var autoInc = uint32(0)

type Record struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"'`
}

func ListNames() []Record {
	var records []Record
	for k, v := range m {
		records = append(records, Record{
			Id:   k,
			Name: v,
		})
	}
	return records
}

func AddName(name string) bool {
	atomic.AddUint32(&autoInc, 1)
	m[autoInc] = name
	return true
}

func DeleteName(id uint32) bool {
	_, ok := m[id]
	delete(m, id)
	return ok
}
