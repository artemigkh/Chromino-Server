package chrominoserver

import (
	"github.com/gin-gonic/gin"
)

func createGameHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:4200")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		game.NewGame()
	}
}
