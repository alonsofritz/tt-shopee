package persistence

import (
	"errors"

	"github.com/alonsofritz/tt-shopee/internal/domain/model"
	"github.com/alonsofritz/tt-shopee/internal/domain/repository"
)

type UserRepoMem struct {
	data map[int]model.User
}

func NewUserRepoMem() repository.UserRepository {
	return &UserRepoMem{
		data: map[int]model.User{
			1: {ID: 1, Name: ""},
		},
	}
}

func (r *UserRepoMem) FindByID(id int) (*model.User, error) {
	if user, ok := r.data[id]; ok {
		return &user, nil
	}
	return nil, errors.New("show not found")
}

func (r *UserRepoMem) Exists(id int) (bool, error) {
	if _, ok := r.data[id]; ok {
		return true, nil
	}
	return false, errors.New("show not found")
}
