package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	Item{
		Id:   1,
		Name: "Laptop",
	},
	Item{
		Id:   2,
		Name: "Mobile",
	},
	Item{
		Id:   3,
		Name: "Monitor",
	},
}

func main() {
	r := gin.Default()

	r.GET("/ping", Ping)
	r.GET("/items", GetItems)
	r.GET("/item/:id", GetItem)
	r.POST("/item", CreateItem)
	r.PUT("/item/:id", UpdateItem)
	r.DELETE("/item/:id", DeleteItem)

	r.Run("localhost:8080")
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetItems(c *gin.Context) {
	fmt.Println("All the items are being returned to the client")
	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")

	fmt.Printf("Item with id %v is being fetched\n", id)

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "seems you have given an incorrect id",
		})

		return
	}

	for _, item := range items {
		if item.Id == int64(intId) {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "the item you are looking for is not available",
	})
}

func CreateItem(c *gin.Context) {
	fmt.Println("Create item API is being called!")
}

func UpdateItem(c *gin.Context) {
	fmt.Println("Create item API is being called!")
}

func DeleteItem(c *gin.Context) {
	fmt.Println("Delete item API is being called!")
}
