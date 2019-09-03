package chrominoserver

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func drawFromStockHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:4200")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		playerIdString := ctx.DefaultQuery("playerId", "0")
		playerId, err := strconv.Atoi(playerIdString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "IdMustBeInteger"})
		}

		countString := ctx.DefaultQuery("count", "1")
		count, err := strconv.Atoi(countString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "IdMustBeInteger"})
		}

		log.Printf("Receieved request for %v pieces from player with id %v", count, playerId)
		if count > 1 {
			ctx.JSON(http.StatusOK, game.DrawPieces(count, playerId))
		} else {
			ctx.JSON(http.StatusOK, game.DrawPiece(playerId, false))
		}
	}
}
