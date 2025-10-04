package persistence

import (
	"errors"

	"github.com/alonsofritz/tt-shopee/internal/domain/model"
	"github.com/alonsofritz/tt-shopee/internal/domain/repository"
)

type ShowRepoMem struct {
	data map[string]model.Show
}

func NewShowRepoMem() repository.ShowRepository {
	return &ShowRepoMem{
		data: map[string]model.Show{
			"230920": {ID: "230920", Description: ""},
		},
	}
}

func (r *ShowRepoMem) FindByID(id string) (*model.Show, error) {
	if show, ok := r.data[id]; ok {
		return &show, nil
	}
	return nil, errors.New("show not found")
}

func (r *ShowRepoMem) Exists(id string) (bool, error) {
	if _, ok := r.data[id]; ok {
		return true, nil
	}
	return false, errors.New("show not found")
}
