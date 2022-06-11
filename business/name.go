package business

import (
	"sync/atomic"
)

type CreateNameIn struct {
	Name string `json:"name"`
}

type Record struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"'`
}

type CreateNameOut struct {
	Success bool     `json:"success"`
	Records []Record `json:"records"`
}

var m = map[uint32]string{}
var autoInc = uint32(1)

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

func CreateName(in *CreateNameIn) (out CreateNameOut) {
	atomic.AddUint32(&autoInc, 1)
	m[autoInc] = in.Name

	out.Success = true
	out.Records = ListNames()

	return
}

type DeleteNameIn struct {
	Id uint32 `json:"id"`
}

type DeleteNameOut struct {
	Success bool     `json:"success"`
	Records []Record `json:"records"`
}

func DeleteName(in *DeleteNameIn) (out DeleteNameOut) {
	_, ok := m[in.Id]
	delete(m, in.Id)

	out.Success = ok
	out.Records = ListNames()

	return
}
