package order_request

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
	order_body "vcs-kafka-learning-go-client/RequestBody"

	api "vcs-kafka-learning-go-client/API"
)

const (
	MAX_QUANTITY       = 15
	MAX_MULTIPLE_ORDER = 100
)

var (
	internalItemsList []order_body.Item
	globalRandom      = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func SendARandomOrder() {
	// Seed the random number generator
	globalRandom.Seed(time.Now().UnixNano())

	// Generate a random index within the range of the items slice length
	randomIndex := globalRandom.Intn(len(internalItemsList))

	// Select a random item
	randomItem := internalItemsList[randomIndex]

	// Generate a random quantity between 1 and 15 (MAX_QUANTITY)
	randomQuantity := globalRandom.Intn(MAX_QUANTITY) + 1

	// Send order request
	orderRequest := order_body.CreateOrderRequest{
		Item: order_body.Item{
			ID:       randomItem.ID,
			Name:     randomItem.Name,
			Quantity: randomQuantity,
		},
	}
	err := orderRequest.Send()
	if err != nil {
		fmt.Println("Error sending order request:", err)
		return
	}
	fmt.Println("Order sent successfully!")
}
func SendMultipleRandomOrders() {
	// fmt.Printf("Enter the number of orders: ")
	// var numberOfOrders int
	// fmt.Scanln(&numberOfOrders)

	// for i := 0; i < numberOfOrders; i++ {
	// 	SendARandomOrder()
	// }
	// fmt.Println("All orders sent successfully!")
	for i := 0; i < MAX_MULTIPLE_ORDER; i++ {
		SendARandomOrder()
	}
	fmt.Printf("%d orders sent successfully!", MAX_MULTIPLE_ORDER)
}
func FetchItems(items *[]order_body.Item) {

	response, err := http.Get(api.Endpoint + "/items")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal response body into OrderItem slice
	err = order_body.UnmarshalFromResponse(responseBody, items)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return
	}

	// Display items
	// order_body.PrintList(items)
	setItemList(items)
}
func setItemList(items *[]order_body.Item) {
	internalItemsList = *items
}
