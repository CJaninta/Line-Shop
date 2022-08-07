package controllers

import (
	"assignment/models"
	"assignment/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {

	var book models.BookReq
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	err := services.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetBook(ctx *gin.Context) {

	books, err := services.GetBook()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func SearchBook(ctx *gin.Context) {

	param := ctx.Query("query")
	if strings.Compare(param, "query") == 0 {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": "Param value is invalid.",
		})
		return
	}

	searchBook, err := services.GetSearchBook(param)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, searchBook)
}
