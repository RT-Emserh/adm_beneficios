package controller

import (
	"net/http"

	"beneficios.com/internal/entities"
	"beneficios.com/internal/service/beneficios"
	"github.com/gin-gonic/gin"
)

func PostBeneficios(ctx *gin.Context) {
	var beneficio entities.Beneficios

	if err := ctx.ShouldBindJSON(&beneficio); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := beneficios.PostBeneficios(&beneficio); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Beneficio criado com sucesso!"})
}
