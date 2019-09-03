package chrominoserver

import (
	"Chromino-Server/chrominoserver/gamemodel"
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var game = gamemodel.CreateChrominoGame()

func StartServing() {
	app := gin.Default()
	app.Use(cors.Default())

	app.Any("/piecePlacement", piecePlacementHandler())
	app.Any("/stateChange", stateChangeHandler())
	app.GET("/drawFromStock", drawFromStockHandler())
	app.POST("/createNewGame", createGameHandler())

	go sendPiecePlacementEvents()
	go sendStateChanges()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		println(">")
		for scanner.Scan() {
			in := scanner.Text()
			println(">")

			if in == "add" {
				println("adding")
				broadcastPiecePlaced <- &PiecePlacement{
					PlayerId: 1,
					Colors:   []string{"RED", "BLUE", "BLUE"},
					Rotation: 90,
					X:        -3,
					Y:        3,
				}
			}
		}
	}()

	if err := app.Run("localhost:8844"); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
