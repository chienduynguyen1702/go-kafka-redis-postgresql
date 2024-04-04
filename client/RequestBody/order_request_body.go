package order_body

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	api "vcs-kafka-learning-go-client/API"
)

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type CreateOrderRequest struct {
	Item
}

func (orderRequest *CreateOrderRequest) Send() error {
	var orderItem Item
	orderItem.Name = orderRequest.Name
	orderItem.Quantity = orderRequest.Quantity
	requestBody, err := json.Marshal(orderItem)
	if err != nil {
		return fmt.Errorf("error marshalling order request: %w", err)
	}
	fmt.Printf("Request: %s\n", string(requestBody))
	response, err := http.Post(api.Endpoint+"/orders", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error making POST request: %w", err)
	}
	defer response.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	fmt.Println("Response:", string(responseBody))
	fmt.Println("")
	return nil
}

func UnmarshalFromResponse(responseBody []byte, items *[]Item) error {
	err := json.Unmarshal(responseBody, items)
	if err != nil {
		return fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return nil
}
func PrintList(items *[]Item) {
	fmt.Printf("List of items:\n")
	fmt.Printf("%5s|%10s|%10s\n", "ID  ", "Name   ", "Quantity ")
	fmt.Printf("-----|----------|----------\n")
	for _, item := range *items {
		item.PrintEach()
	}
	fmt.Printf("-----|----------|----------\n\n")
}
func (item *Item) PrintEach() {
	fmt.Printf("%4d |%-10s| %-9d\n", item.ID, item.Name, item.Quantity)
}
