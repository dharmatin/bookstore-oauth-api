package db

import (
	"github.com/dharmatin/bookstore-oauth-api/src/clients/cassandra"
	"github.com/dharmatin/bookstore-oauth-api/src/domain/access_token"
	"github.com/dharmatin/bookstore-oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken   = "select access_token, user_id, client_id, expires from access_tokens where access_token=?;"
	queryInsertToken      = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpiration = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct{}

func New() DbRepository {
	return &dbRepository{}
}

func (repo *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestError) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (repo *dbRepository) Create(at access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryInsertToken, at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (repo *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryUpdateExpiration, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
