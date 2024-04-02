package models

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"string"`
	Quantity int    `json:"quantity"`
}

var FoodItems = []Item{
	{
		ID:       1,
		Name:     "Apple",
		Quantity: 58,
	},
	{
		ID:       2,
		Name:     "Banana",
		Quantity: 30,
	},
	{
		ID:       3,
		Name:     "Orange",
		Quantity: 13,
	},
	{
		ID:       4,
		Name:     "Pineapple",
		Quantity: 80,
	},
	{
		ID:       5,
		Name:     "Gum",
		Quantity: 22,
	},
	{
		ID:       6,
		Name:     "Chocolate",
		Quantity: 14,
	},
}
