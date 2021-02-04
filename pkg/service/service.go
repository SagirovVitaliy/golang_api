package service

import (
	rest "github.com/SagirovVitaliy/rest_api"
	"github.com/SagirovVitaliy/rest_api/pkg/repository"
)

// Authorization ...
type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// List ...
type List interface {
	Create(userID int, list rest.List) (int, error)
}

// Item ... 
type Item interface {

}

// Service ...
type Service struct {
	Authorization
	List
	Item
}

// NewService ...
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}