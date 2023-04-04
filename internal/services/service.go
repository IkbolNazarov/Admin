package services

import (
	"admin/internal/repository"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

// func (s *Services) GetUsers() ([]*models.Pagination, error) {
// 	return s.Repository.GetData(context)
// }

// func (s *Services) AddData(context.Context) (*models.UserInfo, error) {
// 	return s.Repository.AddData(context)
// }
