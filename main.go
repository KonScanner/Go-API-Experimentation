package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
	MaxQuantity int `json:"maxQuantity"`
}

var books = []book{
	{ID: "1", Title: "The Alchemist", Author: "Paulo Coelho", Quantity: 10, MaxQuantity: 10},
	{ID: "2", Title: "The Kite Runner", Author: "Khaled Hosseini", Quantity: 5, MaxQuantity: 5},
	{ID: "3", Title: "The Da Vinci Code", Author: "Dan Brown", Quantity: 7, MaxQuantity: 7},
	{ID: "4", Title: "The Godfather", Author: "Mario Puzo", Quantity: 3, MaxQuantity: 3},
}

func getBooks(c* gin.Context){
	c.IndentedJSON(http.StatusOK, books)

}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return 
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book,error){
	for i,b := range books{
		if b.ID == id{
			return &books[i],nil
		}
	}

	return nil, errors.New("Book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


func checkID(c* gin.Context) (*book){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id in query"})
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
	return book
}

func checkoutBooks(c* gin.Context){
	book := checkID(c)

	if book.Quantity <= 0{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available"})
		return 
	}
	book.Quantity -=  1
	c.IndentedJSON(http.StatusOK, book)

}

func returnBooks(c* gin.Context){
	book := checkID(c)

	if book.Quantity >= book.MaxQuantity{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "You cannot return a book that is not checked out"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func main() {
	router:= gin.Default()
	router.GET("/books",getBooks)
	router.POST("/books",createBook)
	router.GET("/books/:id",bookById)
	router.PATCH("/checkout",checkoutBooks)
	router.PATCH("/return",returnBooks)
	router.Run("localhost:8080")
	
}