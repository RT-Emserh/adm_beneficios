package main

import (
	"beneficios.com/cmd/api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := gin.Default()
	router.Router(ctx)
	ctx.Run(":3013")
}
