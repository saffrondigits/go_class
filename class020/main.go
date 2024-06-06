package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saffrondigits/go_class/class020/database"
)

type Item struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Handler struct {
	dbClient *sql.DB
}

func main() {
	dbCli, err := database.DbCon()
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = dbCli
	fmt.Printf("Database connection established")
	r := gin.Default()

	handler := &Handler{
		dbClient: dbCli,
	}

	r.GET("/ping", handler.Ping)
	r.GET("/items", handler.GetItems)
	// r.GET("/item/:id", handler.GetItem)
	r.POST("/item", handler.CreateItem)
	// r.PUT("/item/:id", handler.UpdateItem)
	// r.DELETE("/item/:id", handler.DeleteItem)

	r.Run("localhost:8080")
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) GetItems(c *gin.Context) {
	fmt.Println("All the items are being returned to the client")
	rows, err := h.dbClient.Query("SELECT * FROM items;")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "cannot find the items",
		})
		return
	}
	var items []Item
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.Id, &item.Name)
		if err != nil {
			fmt.Println(err)
		}

		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

// func (h *Handler) GetItem(c *gin.Context) {
// 	id := c.Param("id")

// 	fmt.Printf("Item with id %v is being fetched\n", id)

// 	intId, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "seems you have given an incorrect id",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{
// 		"message": "the item you are looking for is not available",
// 	})
// }

func (h *Handler) CreateItem(c *gin.Context) {
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

	res, err := h.dbClient.Exec("INSERT INTO items (name) VALUES ($1)", newItem.Name)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "database error",
		})
	}

	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "database error",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "the item has been added successfully",
	})
}

// func (h *Handler) UpdateItem(c *gin.Context) {
// 	fmt.Println("Create item API is being called!")

// 	id := c.Param("id")
// 	intId, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "seems you have given an incorrect id",
// 		})

// 		return
// 	}

// 	var newItem Item
// 	err = c.ShouldBindJSON(&newItem)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "incorrect json request",
// 		})
// 		return
// 	}

// 	for i, item := range items {
// 		if item.Id == int64(intId) {
// 			items[i].Name = newItem.Name
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "the item has been updated successfully",
// 			})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{
// 		"message": "the item does not exist",
// 	})
// }

// func (h *Handler) DeleteItem(c *gin.Context) {
// 	fmt.Println("Delete item API is being called!")
// 	id := c.Param("id")
// 	intId, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "seems you have given an incorrect id",
// 		})

// 		return
// 	}

// 	for i, item := range items {
// 		if item.Id == int64(intId) {
// 			items = append(items[:i], items[i+1:]...)
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "the item has been deleted successfully",
// 			})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{
// 		"message": "doesn't exist",
// 	})
// }

// Programming
// System and services
// Database design and query
// Scripting
