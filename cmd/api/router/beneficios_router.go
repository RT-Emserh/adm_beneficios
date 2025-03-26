package router

import (
	"beneficios.com/cmd/api/controller"
	"github.com/gin-gonic/gin"
)

func Router(eng *gin.Engine) {
	eng.GET("/health", controller.Health)
}
