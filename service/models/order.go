package models

type Order struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
type CreateOrderRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
