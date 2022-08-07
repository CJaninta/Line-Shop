package controllers

import (
	"assignment/models"
	"assignment/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePromotion(ctx *gin.Context) {

	var promotion models.PromotionReq
	if err := ctx.ShouldBindJSON(&promotion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	err := services.CreatePromotion(&promotion)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetPromotion(ctx *gin.Context) {

	promotions, err := services.GetPromotion()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, promotions)
}
