package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID string     `json:"id"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	Quantity int	`json:"quantity"`
}


var books = []book{
	{ID:"1",Title: "In Search of Lost Time",Author: "Stephen Omondi",Quantity: 2},
	{ID:"2",Title: "The Great Story",Author: "Stephen Ondeyo",Quantity: 5},
	{ID:"3",Title: "Art of Seduction",Author: "Stephen Ochieng",Quantity: 6},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}


func bookById(c *gin.Context){
	id := c.Param("id")
	book,err := getBookById(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK,book)
}



func getBookById(id string)(*book, error){
	for i, b := range books{
		if b.ID == id{
			return &books[i],nil
		}
	}
	return nil,errors.New("Book not found")
}

func createBook(c *gin.Context){
	var newBook book

	if err := c.BindJSON(&newBook); err != nil{
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated,newBook)
}

func checkoutBook(c *gin.Context){
	id,ok := c.GetQuery("id")
	 if !ok{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Missing id querry parameter."})
		return
	 }
	 book,err := getBookById(id)
	 if err != nil {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Missing id querry parameter."})
	 }
	 if book.Quantity <= 0{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Book not available"})
	 }
	 book.Quantity -= 1
	 c.IndentedJSON(http.StatusOK,book)
}


func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.POST("books",createBook)
	router.GET("/books/:id",bookById)
	router.PATCH("checkout",checkoutBook)
	router.Run("localhost:8081")
}