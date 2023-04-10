package entity

const (
	Dough     = "d"
	Cheese    = "c"
	Additives = "i"
)

type Ingredient struct {
	ID        int64   `db:"id"`
	TypeID    int64   `db:"type_id"`
	TypeTitle string  `db:"type_title"`
	TypeCode  string  `db:"type_code"`
	Title     string  `db:"title"`
	Price     float64 `db:"price"`
}

type IngredientType struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
	Code  string `db:"code"`
}

type IngredientsGroupByType struct {
	Dough     []Ingredient
	Cheese    []Ingredient
	Additives []Ingredient
}

type Product struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type DishMask struct {
	Dough     int
	Cheese    int
	Additives int
}

type Dish struct {
	Products []Product `json:"products"`
	Price    float64   `json:"price"`
}

type Menu []Dish
