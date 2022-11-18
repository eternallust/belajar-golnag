package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"pustaka-api/book"

	"github.com/go-playground/validator/v10"
)


func RootHandler(c *gin.Context){
    c.JSON(http.StatusOK, gin.H {
        "name": "Rizky Abdulrachman",
        "bio": "Software engineer",
    })
}

func HelloHandler(c *gin.Context){
    c.JSON(http.StatusOK, gin.H {
        "title": "hello world",
        "subtitle": "belajar golang",
    })
}

// localhost:8888/books/halah
func BooksHandler(c *gin.Context){
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H {
        "id": id,
    })
}

// localhost:8888/query?title=halah
func QueryHandler(c *gin.Context){
    title := c.Query("title")
    price := c.Query("price")
    c.JSON(http.StatusOK, gin.H {
        "title": title,
        "price": price,
    })
}

func PostBooksHandler(c *gin.Context){

    var bookInput book.BookRequest;
    
    err := c.ShouldBindJSON(&bookInput);

    if err != nil {
        errorMessages := []string{}
        var ve validator.ValidationErrors
        if errors.As(err, &ve){
            for _, e:= range err.(validator.ValidationErrors){
                errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
                errorMessages = append(errorMessages, errorMessage)    
            }
            c.JSON(http.StatusBadRequest, gin.H{
                "error": errorMessages,
            })
            return
        }
        
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("%v", err),
        })
        return
    }
    
    c.JSON(http.StatusOK, gin.H {
        "title": bookInput.Title,
        "price": bookInput.Price,
    })
}