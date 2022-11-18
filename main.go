package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=pustaka-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	

	if err != nil {
		log.Fatal("DB connection error")
	}

	fmt.Println("DB connected")
	
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	

	bookRequest:= book.BookRequest {
		Title: "buku baru",
		Price: "10000",
	}

	bookService.Create(bookRequest)


	router := gin.Default()

    v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
    v1.GET("/books/:id", handler.BooksHandler)
    v1.GET("/query", handler.QueryHandler)

    v1.POST("/books", handler.PostBooksHandler)
    
	router.Run(":8888")
}



