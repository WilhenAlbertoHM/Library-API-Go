package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This is a simple library API, where we can check all books available in the library,
// check book by ID, and check in and out books.

// Fields start with uppercase letters to make them exported/public fields,
// for other models to view them.
// We also convert them into JSON and viceversa when we use the API
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// List of books.
var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// Handle the route of getting all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Helper function to get a book by ID.
func getBookByID(id string) (*book, error) {
	// For each book in the list of books,
	// check if the ID matches the ID we are looking for.
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// Handle the route of getting a book by ID.
func bookByID(c *gin.Context) {
	// Get the ID from the URL.
	id := c.Param("id")
	book, err := getBookByID(id)

	// If the book is not found, return an error.
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	// Else, return the book.
	c.IndentedJSON(http.StatusOK, book)
}

// Handle the route of creating a new book.
func createBook(c *gin.Context) {
	// Get data for book
	var newBook book

	// Bind the request data to the newBook variable.
	// If there is an error, return.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the list of books.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// Handle the route of checking out a book.
func checkOutBook(c *gin.Context) {
	// Get the ID from the URL.
	id, ok := c.GetQuery("id")

	// If the ID is not found, return an error.
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	// Get the book by ID.
	book, err := getBookByID(id)

	// If the book is not found, return an error.
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	// If the quantity of the book is 0, return an error.
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book is not available"})
		return
	}

	// Decrement the quantity of the book by 1 and return the book.
	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)
}

// Handle the route of checking in a book.
func checkInBook(c *gin.Context) {
	// Get the ID from the URL.
	id, ok := c.GetQuery("id")

	// If the ID is not found, return an error.
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	// Get the book by ID.
	book, err := getBookByID(id)

	// If the book is not found, return an error.
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	// Increment the quantity of the book by 1 and return the book.
	book.Quantity++
	c.IndentedJSON(http.StatusOK, book)
}

// Main function.
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookByID)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/checkin", checkInBook)
	router.Run("localhost:8080")
}
