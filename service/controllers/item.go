package controllers

import (
	"fmt"
	"net/http"
	"vcs-kafka-learning-go-service/models"

	"github.com/gin-gonic/gin"
)

//	GetItem godoc
//
// @Summary Get all items and its quantity
// @Schemes
// @Description  Get all items and its quantity
// @Tags Items
// @Accept json
// @Produce json
// @Success 200 {araray} models.Items "An array of items with their quantities."
// @Router /api/items [get]
func GetItem(c *gin.Context) {

	// Get items from Redis along with their quantities
	var items []models.Item
	for _, item := range models.FoodItems {
		key := fmt.Sprintf("food:%s:quantity", item.Name)
		quantity, err := RedisClient.HGet(ctxBackground, key, "quantity").Int()
		if err != nil {
			// Handle error (e.g., log it or return an error response)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item quantities from Redis"})
			return
		}
		// Append item with quantity to the list
		items = append(items, models.Item{
			ID:       item.ID,
			Name:     item.Name,
			Quantity: quantity,
		})
	}

	// Return items with quantities in JSON response
	c.JSON(http.StatusOK, items)

}

// ResetItemQuantity godoc
//
// @Summary Reset quantity of items as start
// @Schemes
// @Description  Reset quantity of items as start
// @Tags Items
// @Accept json
// @Produce json
// @Success 200 {string}  string "Items have been reset successfully."
// @Router /api/items [post]
func ResetItemQuantity(c *gin.Context) {
	var foodItems = models.FoodItems
	var err error
	for _, item := range foodItems {
		key := fmt.Sprintf("food:%s:quantity", item.Name)
		err = RedisClient.HSet(ctxBackground, key, "quantity", item.Quantity).Err()
		if err != nil {
			fmt.Println("Error setting quantity for", item.Name, ":", err)
		}
	}
	if err == nil {
		c.String(http.StatusOK, "Items have been reset successfully.")
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "An error occurred while trying to reset items"})
	}
}
