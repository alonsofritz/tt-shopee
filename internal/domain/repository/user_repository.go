package repository

import "github.com/alonsofritz/tt-shopee/internal/domain/model"

type UserRepository interface {
	FindByID(id int) (*model.User, error)
}
