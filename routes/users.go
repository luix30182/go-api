package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mario-mtz.com/rest-api/models"
)

func signup(context *gin.Context) {
	var user models.USER
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created"})

}
