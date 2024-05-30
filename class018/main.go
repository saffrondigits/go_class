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
	var newItem Item

	err := c.ShouldBindJSON(&newItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "seems data you are sending is incorrect",
		})

		return
	}

	if newItem.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "seems data you are sending is incorrect",
		})

		return
	}

	for _, v := range items {
		if v.Name == newItem.Name {
			c.JSON(http.StatusConflict, gin.H{
				"message": "the item already exists",
			})
			return
		}
	}

	newItem.Id = int64(len(items) + 1)

	items = append(items, newItem)
	c.JSON(http.StatusCreated, gin.H{
		"message": "the item has been added successfully",
	})
}

func UpdateItem(c *gin.Context) {
	fmt.Println("Create item API is being called!")

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "seems you have given an incorrect id",
		})

		return
	}

	var newItem Item
	err = c.ShouldBindJSON(&newItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect json request",
		})
		return
	}

	for i, item := range items {
		if item.Id == int64(intId) {
			items[i].Name = newItem.Name
			c.JSON(http.StatusOK, gin.H{
				"message": "the item has been updated successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "the item does not exist",
	})
}

func DeleteItem(c *gin.Context) {
	fmt.Println("Delete item API is being called!")
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "seems you have given an incorrect id",
		})

		return
	}

	for i, item := range items {
		if item.Id == int64(intId) {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "the item has been deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "doesn't exist",
	})
}
