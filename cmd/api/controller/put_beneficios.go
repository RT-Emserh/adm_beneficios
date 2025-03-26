package controller

import (
	"net/http"

	"beneficios.com/internal/entities"
	"beneficios.com/internal/service/beneficios"
	"github.com/gin-gonic/gin"
)

func PutBeneficios(ctx *gin.Context) {
	var beneficio entities.Beneficios
	id := ctx.Param("ID")

	if err := ctx.ShouldBindJSON(&beneficio); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"mensagem erro": "Dados inválidos"})
		return
	}

	if err := beneficios.PutBeneficios(id, &beneficio); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"mensagem erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensagem": "Benefício atualizado com sucesso"})
}
