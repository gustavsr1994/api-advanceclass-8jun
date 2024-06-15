package controllers

import (
	"example/api-advance-class/config"
	"example/api-advance-class/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func All(c *gin.Context) {
	var product []models.Product
	statusErr := false
	message := "Success"
	config.DB.Find(&product)
	if len(product) == 0 {
		statusErr = true
		message = "No Data"
	}
	c.JSON(http.StatusOK, gin.H{"data": product, "error": statusErr, "message": message})
}

func Index(c *gin.Context) {
	var product models.Product
	statusErr := false
	message := "Success"

	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			statusErr = true
			message = "Data not found"
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"data": product, "error": statusErr, "message": message})
			return
		default:
			statusErr = true
			message = err.Error()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": product, "error": statusErr, "message": message})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": product, "error": statusErr, "message": message})
}

func Create(c *gin.Context) {
	var product models.Product
	statusErr := false
	message := "Success"
	if err := c.ShouldBindJSON(&product); err != nil {
		statusErr = true
		message = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data": product, "error": statusErr, "message": message})
		return
	}

	if config.DB.Create(&product).RowsAffected == 0 {
		statusErr = true
		message = "Tidak dapat insert new data"
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data": product, "error": statusErr, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product, "error": statusErr, "message": message})
}

func Update(c *gin.Context) {
	var product models.Product
	statusErr := false
	message := "Success"
	if err := c.ShouldBindJSON(&product); err != nil {
		statusErr = true
		message = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data": product, "error": statusErr, "message": message})
		return
	}

	id := product.Id
	if config.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		statusErr = true
		message = "Tidak dapat update data"
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data": product, "error": statusErr, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product, "error": statusErr, "message": message})
}

func Delete(c *gin.Context) {
	var product models.Product
	statusErr := false
	message := "Success"

	id := c.Param("id")
	if config.DB.Delete(&product, id).RowsAffected == 0 {
		statusErr = true
		message = "Tidak dapat delete data"
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"data": product, "error": statusErr, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product, "error": statusErr, "message": message})
}
