package main

import (
	"github.com/eluamous-droid/godo/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getAllItems", database.GetAllItems)
	router.GET("/getAllItemsInGroup", database.GetAllItemsInGroup)
	router.POST("/addItem", database.AddItem)
	router.PUT("/deleteItem", database.DeleteItem)
	router.Run("localhost:8080")
}
