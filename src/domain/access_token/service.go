package access_token

import (
	"strings"

	"github.com/MohamedNazir/go_bookstore_oauth_api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {

	accesstokenid := strings.TrimSpace(accessTokenID)
	if len(accesstokenid) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetByID(accessTokenID)

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {

	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
