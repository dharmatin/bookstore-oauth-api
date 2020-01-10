package access_token

import (
	"strings"

	"github.com/dharmatin/bookstore-oauth-api/src/domain/users"
	"github.com/dharmatin/bookstore-oauth-api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestError)
	Create(AccessTokenRequest) (*AccessToken, *errors.RestError)
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type service struct {
	dbRepo       Repository
	restUserRepo RestUserRepository
}

func NewService(restUserRepo RestUserRepository, repo Repository) Service {
	return &service{
		dbRepo:       repo,
		restUserRepo: restUserRepo,
	}
}

func (s *service) GetByID(id string) (*AccessToken, *errors.RestError) {
	accessTokenId := strings.TrimSpace(id)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token")
	}
	accessToken, err := s.dbRepo.GetByID(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request AccessTokenRequest) (*AccessToken, *errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	at := GetNewAccessToken(user.Id)
	at.Generate()
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
