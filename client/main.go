package main

import (
	"fmt"
	"os"
	order_request "vcs-kafka-learning-go-client/Request"
	order_body "vcs-kafka-learning-go-client/RequestBody"
)

func main() {
	// menu to select the operation
	menu()

}
func menu() {
	var items []order_body.Item
	var choice int
	for {
		fmt.Println("Select the operation")
		fmt.Println("1. List all items")
		fmt.Println("2. Make one random order")
		fmt.Println("3. Make multiple random orders")
		fmt.Println("0. Exit")
		fmt.Print("Enter the operation number: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch choice {
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		case 1:
			order_request.ListAllItems(&items)
		case 2:
			order_request.SendARandomOrder()
		case 3:
			order_request.SendMultipleRandomOrders()
		default:
			fmt.Println("Invalid operation number. Please try again.")
		}
	}

}
