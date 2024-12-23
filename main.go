package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
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

func createBook(c *gin.Context){
	var newBook book
}

func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.Run("localhost:8081")
}