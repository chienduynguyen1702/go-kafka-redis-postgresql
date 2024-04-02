package models

type Response struct {
	OrderID   int    `json:"id"`
	IsSucceed bool   `json:"is_succeed"`
	Message   string `json:"message"`
}
