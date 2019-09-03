package gamemodel

type StateDisplay struct {
	DeckSize int              `json:"deckSize"`
	Players  []*PlayerDisplay `json:"players"`
}
