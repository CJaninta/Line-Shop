package controllers

import (
	"assignment/models"
	"assignment/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {

	var order models.OrderReq
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	err := services.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func CheckoutOrder(ctx *gin.Context) {
	var orderId models.OrderCheckoutReq
	if err := ctx.ShouldBindJSON(&orderId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	err := services.CheckoutOrder(orderId.Id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func Summary(ctx *gin.Context) {
	orderSummary, err := services.Summary()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, orderSummary)
}
