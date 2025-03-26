package controller

import (
	"net/http"

	"beneficios.com/internal/service/beneficios"
	"github.com/gin-gonic/gin"
)

func DeleteBeneficios(ctx *gin.Context) {
	id := ctx.Param("id")

	err := beneficios.DeleteBeneficios(id)
	if err != nil {
		if err.Error() == "Beneficio não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"mensagem erro": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"mensagem erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Beneficio excluído com sucesso!",
	})
}
