package repository

import (
	"fmt"

	rest "github.com/SagirovVitaliy/rest_api"
	"github.com/jmoiron/sqlx"
)

// AuthPostgres ...
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres ...
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser ...
func (r *AuthPostgres) CreateUser(user rest.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}


// GetUser ...
func (r *AuthPostgres) GetUser(username, password string) (rest.User, error) {
	var user rest.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}