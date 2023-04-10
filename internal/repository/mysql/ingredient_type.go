package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"goulash-menu-api/internal/entity"
	"strconv"
	"strings"
)

type IngredientType struct {
	db *sqlx.DB
}

func NewIngredientType(db *sqlx.DB) *IngredientType {
	return &IngredientType{db: db}
}

func (r *IngredientType) GetAll() ([]entity.IngredientType, error) {
	var ingredientTypes []entity.IngredientType

	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM %s", IngredientTypeTable))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ingredientType entity.IngredientType
		err = rows.StructScan(&ingredientType)
		if err != nil {
			return nil, err
		}
		ingredientTypes = append(ingredientTypes, ingredientType)
	}

	return ingredientTypes, nil
}

func (r *IngredientType) GetByID(ID int64) (*entity.IngredientType, error) {
	var ingredientType entity.IngredientType

	row := r.db.QueryRowx(fmt.Sprintf("SELECT * FROM %s WHERE id = ?", IngredientTable), ID)

	err := row.StructScan(&ingredientType)
	if err != nil {
		return nil, err
	}

	return &ingredientType, nil
}

func (r *IngredientType) GetAllByIDs(typeIDs []int64) ([]entity.IngredientType, error) {
	var ingredientTypes []entity.IngredientType

	var textIDs []string
	for _, v := range typeIDs {
		textIDs = append(textIDs, strconv.Itoa(int(v)))
	}
	IDs := strings.Join(textIDs, ",")

	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM %s WHERE id in (%s)", IngredientTypeTable, IDs))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ingredientType entity.IngredientType
		err = rows.StructScan(&ingredientType)
		if err != nil {
			return nil, err
		}
		ingredientTypes = append(ingredientTypes, ingredientType)
	}

	return ingredientTypes, nil
}
