package service

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"goulash-menu-api/internal/entity"
	"testing"
)

var ingredients = entity.IngredientsGroupByType{
	Dough: []entity.Ingredient{
		{
			ID:        1,
			TypeID:    1,
			TypeTitle: "Тесто",
			TypeCode:  "d",
			Title:     "Тонкое тесто",
			Price:     100,
		},
		{
			ID:        2,
			TypeID:    1,
			TypeTitle: "Тесто",
			TypeCode:  "d",
			Title:     "Пышное тесто",
			Price:     110,
		},
		{
			ID:        3,
			TypeID:    1,
			TypeTitle: "Тесто",
			TypeCode:  "d",
			Title:     "Ржаное тесто",
			Price:     150,
		},
	},
	Cheese: []entity.Ingredient{
		{
			ID:        4,
			TypeID:    2,
			TypeTitle: "Сыр",
			TypeCode:  "c",
			Title:     "Моцарелла",
			Price:     50,
		},
		{
			ID:        5,
			TypeID:    2,
			TypeTitle: "Сыр",
			TypeCode:  "c",
			Title:     "Рикотта",
			Price:     70,
		},
	},
	Additives: []entity.Ingredient{
		{
			ID:        6,
			TypeID:    3,
			TypeTitle: "Начинка",
			TypeCode:  "i",
			Title:     "Колбаса",
			Price:     30,
		},
		{
			ID:        7,
			TypeID:    3,
			TypeTitle: "Начинка",
			TypeCode:  "i",
			Title:     "Ветчина",
			Price:     35,
		},
		{
			ID:        8,
			TypeID:    3,
			TypeTitle: "Начинка",
			TypeCode:  "i",
			Title:     "Грибы",
			Price:     50,
		},
		{
			ID:        9,
			TypeID:    3,
			TypeTitle: "Начинка",
			TypeCode:  "i",
			Title:     "Томаты",
			Price:     10,
		},
	},
}

func Test_generateMenu(t *testing.T) {
	type args struct {
		dishMask    entity.DishMask
		ingredients entity.IngredientsGroupByType
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				dishMask:    dishMaskFromString(""),
				ingredients: ingredients,
			},
			want: `[]`,
		},
		{
			name: "dc",
			args: args{
				dishMask:    dishMaskFromString("dc"),
				ingredients: ingredients,
			},
			want: `[]`,
		},
		{
			name: "dciiiii",
			args: args{
				dishMask:    dishMaskFromString("dciiiii"),
				ingredients: ingredients,
			},
			want: `[]`,
		},
		{
			name: "ddcci",
			args: args{
				dishMask:    dishMaskFromString("ddcci"),
				ingredients: ingredients,
			},
			want: `[{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"},{"type":"Тесто","value":"Пышное тесто"}],"price":360},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"},{"type":"Тесто","value":"Ржаное тесто"}],"price":400},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"},{"type":"Тесто","value":"Ржаное тесто"}],"price":410},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"},{"type":"Тесто","value":"Пышное тесто"}],"price":365},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"},{"type":"Тесто","value":"Ржаное тесто"}],"price":405},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"},{"type":"Тесто","value":"Ржаное тесто"}],"price":415},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"},{"type":"Тесто","value":"Пышное тесто"}],"price":380},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"},{"type":"Тесто","value":"Ржаное тесто"}],"price":420},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"},{"type":"Тесто","value":"Ржаное тесто"}],"price":430},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"},{"type":"Тесто","value":"Пышное тесто"}],"price":340},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"},{"type":"Тесто","value":"Ржаное тесто"}],"price":380},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"},{"type":"Тесто","value":"Ржаное тесто"}],"price":390}]`,
		},
		{
			name: "dci",
			args: args{
				dishMask:    dishMaskFromString("dci"),
				ingredients: ingredients,
			},
			want: `[{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":180},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":190},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":230},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":200},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":210},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":250},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":185},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":195},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":235},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":205},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":215},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":255},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":200},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":210},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":250},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":220},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":230},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":270},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":160},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":170},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":210},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":180},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":190},{"products":[{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":230}]`,
		},
		{
			name: "dcii",
			args: args{
				dishMask:    dishMaskFromString("dcii"),
				ingredients: ingredients,
			},
			want: `[{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":215},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":225},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":265},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":235},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":245},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":285},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":230},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":240},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":280},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":250},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":260},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":300},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":190},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":200},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":240},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":210},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":220},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":260},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":235},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":245},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":285},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":255},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":265},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":305},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":195},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":205},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":245},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":215},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":225},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":265},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":210},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":220},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":260},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":230},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":240},{"products":[{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":280}]`,
		},
		{
			name: "dciii",
			args: args{
				dishMask:    dishMaskFromString("dciii"),
				ingredients: ingredients,
			},
			want: `[{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":265},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":275},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":315},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":285},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":295},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":335},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":225},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":235},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":275},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":245},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":255},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":295},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":240},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":250},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":290},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":260},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":270},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":310},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":245},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":255},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":295},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":265},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":275},{"products":[{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":315}]`,
		},
		{
			name: "dciiii",
			args: args{
				dishMask:    dishMaskFromString("dciiii"),
				ingredients: ingredients,
			},
			want: `[{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Тонкое тесто"}],"price":275},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Пышное тесто"}],"price":285},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Моцарелла"},{"type":"Тесто","value":"Ржаное тесто"}],"price":325},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Тонкое тесто"}],"price":295},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Пышное тесто"}],"price":305},{"products":[{"type":"Начинка","value":"Колбаса"},{"type":"Начинка","value":"Ветчина"},{"type":"Начинка","value":"Грибы"},{"type":"Начинка","value":"Томаты"},{"type":"Сыр","value":"Рикотта"},{"type":"Тесто","value":"Ржаное тесто"}],"price":345}]`,
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := json.Marshal(generateMenu(tt.args.dishMask, tt.args.ingredients))
			if diff := cmp.Diff(got, []byte(tt.want)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
