package business

import (
	"sveltefiber/model"
)

// TODO: refactor to struct with model dependency

type CreateNameIn struct {
	Name string `json:"name"`
}

type CreateNameOut struct {
	Success bool           `json:"success"`
	Records []model.Record `json:"records"`
}

func CreateName(in *CreateNameIn) (out CreateNameOut) {
	out.Success = model.AddName(in.Name)
	out.Records = model.ListNames()
	return
}

type DeleteNameIn struct {
	Id uint32 `json:"id"`
}

type DeleteNameOut struct {
	Success bool           `json:"success"`
	Records []model.Record `json:"records"`
}

func DeleteName(in *DeleteNameIn) (out DeleteNameOut) {
	out.Success = model.DeleteName(in.Id)
	out.Records = model.ListNames()
	return
}
