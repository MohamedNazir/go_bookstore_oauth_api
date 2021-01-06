package db

import (
	"github.com/MohamedNazir/go_bookstore_oauth_api/src/clients/cassendra"
	"github.com/MohamedNazir/go_bookstore_oauth_api/src/domain/access_token"
	"github.com/MohamedNazir/go_bookstore_oauth_api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "select access_token, user_id, client_id, expires from access_tokens where access_token=?"
	queryCreateAccessToken = "Insert into access_tokens (access_token,user_id, client_id, expires) values (?,?,?,?)"
	queryUpdateExpires     = "update access_token set expires =? where access_tokens = ?"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {

	var result access_token.AccessToken
	if err := cassendra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No records found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {

	if err := cassendra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {

	if err := cassendra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
