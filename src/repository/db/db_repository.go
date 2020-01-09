package db

import (
	"github.com/dharmatin/bookstore-oauth-api/src/domain/access_token"
	"github.com/dharmatin/bookstore-oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct{}

func New() DbRepository {
	return &dbRepository{}
}

func (repo *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("Database not implement yet!")
}
