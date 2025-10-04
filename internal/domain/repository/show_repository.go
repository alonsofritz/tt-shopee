package repository

import "github.com/alonsofritz/tt-shopee/internal/domain/model"

type ShowRepository interface {
	FindByID(id string) (*model.Show, error)
}
