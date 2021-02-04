package repository

import (
	rest "github.com/SagirovVitaliy/rest_api"
	"github.com/jmoiron/sqlx"
)

// Authorization ...
type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GetUser(username, password string) (rest.User, error)
}

// List ...
type List interface {
	Create(userID int, list rest.List) (int, error)
}

// Item ... 
type Item interface {

}

// Repository ...
type Repository struct {
	Authorization
	List
	Item
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}