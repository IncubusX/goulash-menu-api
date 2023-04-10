package repository

import "goulash-menu-api/internal/entity"

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type IIngredient interface {
	GetAll() ([]entity.Ingredient, error)
	GetByID(ID int64) (*entity.Ingredient, error)
	GetByType(typeID int64) ([]entity.Ingredient, error)
	GetAllGroupByType() (entity.IngredientsGroupByType, error)
}
type IIngredientType interface {
	GetAll() ([]entity.IngredientType, error)
	GetByID(ID int64) (*entity.IngredientType, error)
	GetAllByIDs(typeIDs []int64) ([]entity.IngredientType, error)
}
