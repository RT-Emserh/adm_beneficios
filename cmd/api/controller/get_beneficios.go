package controller

import (
	"net/http"

	"beneficios.com/internal/service/beneficios"
	"github.com/gin-gonic/gin"
)

func GetEmserhBeneficios(ctx *gin.Context) {
	beneficios, err := beneficios.GetBeneficios()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"mensagem erro": "Erro ao encontrar os benefícios",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"beneficios": beneficios,
	})
}
