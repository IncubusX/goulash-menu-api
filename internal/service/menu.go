package service

import (
	"goulash-menu-api/internal/entity"
	"goulash-menu-api/internal/repository"
)

type MenuService struct {
	Ingredient repository.IIngredient
}

func NewMenuService(repo repository.IIngredient) *MenuService {
	return &MenuService{Ingredient: repo}
}

func (s *MenuService) GenerateMenu(mask string) (entity.Menu, error) {
	ingredients, err := s.Ingredient.GetAllGroupByType()
	if err != nil {
		return nil, err
	}
	dishMask := dishMaskFromString(mask)
	return generateMenu(dishMask, ingredients), nil
}

func dishMaskFromString(mask string) entity.DishMask {
	var dishMask entity.DishMask
	for _, c := range mask {
		switch string(c) {
		case entity.Dough:
			dishMask.Dough++
		case entity.Cheese:
			dishMask.Cheese++
		case entity.Additives:
			dishMask.Additives++
		}
	}
	return dishMask
}

func generateMenu(dishMask entity.DishMask, ingredients entity.IngredientsGroupByType) entity.Menu {
	menu := generateList(entity.Menu{}, entity.Dish{}, ingredients.Additives, dishMask.Additives)
	listSize := len(menu)

	for _, list := range menu {
		menu = generateList(menu, list, ingredients.Cheese, dishMask.Cheese)
	}
	menu = menu[listSize:]
	listSize = len(menu)

	for _, list := range menu {
		menu = generateList(menu, list, ingredients.Dough, dishMask.Dough)
	}
	menu = menu[listSize:]

	return menu
}

func generateList(menu entity.Menu, current entity.Dish, ingredients []entity.Ingredient, length int) entity.Menu {
	if length <= 0 {
		return menu
	}
	for i, ingredient := range ingredients {
		newItem := entity.Dish{}
		newItem.Products = append(newItem.Products, current.Products...)
		newItem.Price = current.Price

		newItem = addToDish(newItem, ingredient)
		if length > 1 {
			menu = generateList(menu, newItem, ingredients[i+1:], length-1)
		} else {
			menu = append(menu, newItem)
		}
		if len(ingredients) <= length {
			break
		}
	}
	return menu
}

func addToDish(dish entity.Dish, ingredient entity.Ingredient) entity.Dish {
	product := entity.Product{
		Type:  ingredient.TypeTitle,
		Value: ingredient.Title,
	}

	dish.Products = append(dish.Products, product)
	dish.Price += ingredient.Price
	return dish
}
