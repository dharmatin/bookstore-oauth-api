package access_token

import (
	"github.com/dharmatin/bookstore-oauth-api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestError)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(id string) (*AccessToken, *errors.RestError) {
	accessToken, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
