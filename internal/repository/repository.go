package repository

import (
	"github.com/jmoiron/sqlx"
	repository "goulash-menu-api/internal/repository/mysql"
)

type (
	Repository struct {
		Ingredient     IIngredient
		IngredientType IIngredientType
	}
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Ingredient:     repository.NewIngredient(db),
		IngredientType: repository.NewIngredientType(db),
	}
}
