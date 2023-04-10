package service

import "goulash-menu-api/internal/repository"

type Service struct {
	Menu IMenu
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Menu: NewMenuService(repos.Ingredient),
	}
}
