package service

import "goulash-menu-api/internal/entity"

type IMenu interface {
	GenerateMenu(ingredients string) (entity.Menu, error)
}
