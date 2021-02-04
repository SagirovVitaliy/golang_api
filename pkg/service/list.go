package service

import (
	rest "github.com/SagirovVitaliy/rest_api"
	"github.com/SagirovVitaliy/rest_api/pkg/repository"
)

// ListService ...
type ListService struct {
	repo repository.List
}

// NewListService ...
func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

// Create ...
func (s *ListService) Create(userID int, list rest.List) (int, error) {
	return s.repo.Create(userID, list)
}