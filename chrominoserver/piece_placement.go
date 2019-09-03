package chrominoserver

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type PiecePlacement struct {
	PlayerId int      `json:"playerId"`
	Colors   []string `json:"colors"`
	Rotation int      `json:"rotation"`
	X        int      `json:"x"`
	Y        int      `json:"y"`
}

var piecePlacementClients = make(map[*websocket.Conn]bool)
var broadcastPiecePlaced = make(chan *PiecePlacement)

func piecePlacementHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		piecePlacementClients[conn] = true

		for {
			var piecePlacement PiecePlacement
			if err = conn.ReadJSON(&piecePlacement); err != nil {
				log.Error().Msg(err.Error())
			}

			log.Print("received piece placement event", piecePlacement)
			broadcastPiecePlaced <- &piecePlacement
		}
	}
}

func sendPiecePlacementEvents() {
	for {
		val := <-broadcastPiecePlaced
		log.Printf("broadcasting piece placement", val)
		for client := range piecePlacementClients {
			err := client.WriteJSON(val)
			if err != nil {
				log.Printf("Websocket error: %s", err)
			}
		}
	}
}
