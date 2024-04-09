package models

type Response struct {
	Message   string `json:"message"`
	IsSucceed bool   `json:"is_succeed"`
}
