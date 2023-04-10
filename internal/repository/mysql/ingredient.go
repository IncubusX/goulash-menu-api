package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"goulash-menu-api/internal/entity"
)

const (
	IngredientTable     = "ingredient"
	IngredientTypeTable = "ingredient_type"
)

type Ingredient struct {
	db *sqlx.DB
}

func NewIngredient(db *sqlx.DB) *Ingredient {
	return &Ingredient{db: db}
}

func (r *Ingredient) GetAll() ([]entity.Ingredient, error) {
	var ingredients []entity.Ingredient

	rows, err := r.db.Queryx(fmt.Sprintf(
		`SELECT i.id, i.type_id, it.title as type_title, it.code as type_code, i.title, i.price 
					FROM %s AS i 
					INNER JOIN %s AS it ON it.id = i.type_id`, IngredientTable, IngredientTypeTable))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ingredient entity.Ingredient
		err = rows.StructScan(&ingredient)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (r *Ingredient) GetAllGroupByType() (entity.IngredientsGroupByType, error) {
	var ingredientsGroupByType entity.IngredientsGroupByType

	ingredients, err := r.GetAll()
	if err != nil {
		return entity.IngredientsGroupByType{}, err
	}

	for _, ingredient := range ingredients {
		switch ingredient.TypeCode {
		case entity.Dough:
			ingredientsGroupByType.Dough = append(ingredientsGroupByType.Dough, ingredient)
		case entity.Cheese:
			ingredientsGroupByType.Cheese = append(ingredientsGroupByType.Cheese, ingredient)
		case entity.Additives:
			ingredientsGroupByType.Additives = append(ingredientsGroupByType.Additives, ingredient)
		}
	}
	return ingredientsGroupByType, nil
}

func (r *Ingredient) GetByID(ID int64) (*entity.Ingredient, error) {
	var ingredient entity.Ingredient

	row := r.db.QueryRowx(fmt.Sprintf("SELECT * FROM %s WHERE id = ?", IngredientTable), ID)

	err := row.StructScan(&ingredient)
	if err != nil {
		return nil, err
	}

	return &ingredient, nil
}

func (r *Ingredient) GetByType(typeID int64) ([]entity.Ingredient, error) {
	var ingredients []entity.Ingredient

	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM %s WHERE type_id = ?", IngredientTable), typeID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ingredient entity.Ingredient
		err = rows.StructScan(&ingredient)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}
