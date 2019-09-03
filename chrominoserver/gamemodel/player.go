package gamemodel

type Player struct {
	PlayerId   int
	Name       string
	Hand       []ChrominoPiece
	ActiveTurn bool
}

type PlayerDisplay struct {
	PlayerId   int    `json:"playerId"`
	Name       string `json:"name"`
	HandSize   int    `json:"handSize"`
	ActiveTurn bool   `json:"activeTurn"`
}

func (player *Player) display() *PlayerDisplay {
	return &PlayerDisplay{
		PlayerId:   player.PlayerId,
		Name:       player.Name,
		HandSize:   len(player.Hand),
		ActiveTurn: false,
	}
}
