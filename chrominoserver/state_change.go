package chrominoserver

import (
	"Chromino-Server/chrominoserver/gamemodel"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var stateChangeClients = make(map[*websocket.Conn]bool)
var broadcastStateChange = make(chan *gamemodel.StateDisplay)

func stateChangeHandler() gin.HandlerFunc {
	println("Setting game callback")
	game.RegisterStateChangeCallback(func(stateChange gamemodel.StateDisplay) {
		log.Print("received state change", stateChange)
		broadcastStateChange <- &stateChange
	})
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		stateChangeClients[conn] = true
	}
}

func sendStateChanges() {
	for {
		val := <-broadcastStateChange
		log.Printf("broadcasting state change", val)
		for client := range stateChangeClients {
			err := client.WriteJSON(val)
			if err != nil {
				log.Printf("Websocket error: %s", err)
			}
		}
	}
}
