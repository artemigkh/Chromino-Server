package gamemodel

import (
	"fmt"
	"sync"
)

type ChrominoGame struct {
	deck                *chrominoDeck
	stateChangeCallback func(display StateDisplay)
	mutex               sync.Mutex
}

func CreateChrominoGame() *ChrominoGame {
	return &ChrominoGame{
		deck: newDeck(),
	}
}

func (game *ChrominoGame) RegisterStateChangeCallback(
	stateChangeCallback func(display StateDisplay)) {
	game.stateChangeCallback = stateChangeCallback
}

func (game *ChrominoGame) NewGame() {
	game.mutex.Lock()
	defer game.mutex.Unlock()

	game.deck = newDeck()
}

func (game *ChrominoGame) DrawPiece(_ int, multiple bool) [3]ChrominoColor {
	game.mutex.Lock()
	defer game.mutex.Unlock()

	piece, remaining := game.deck.draw()
	fmt.Printf("Drawing piece: [%v %v %v]\n",
		piece.Colors[0], piece.Colors[1], piece.Colors[2])
	fmt.Printf("Remaining pieces: %v\n", remaining)

	if !multiple {
		game.stateChangeCallback(game.generateDisplay())
	}
	return piece.Colors
}

func (game *ChrominoGame) DrawPieces(count int, _ int) [][3]ChrominoColor {

	batch := make([][3]ChrominoColor, 0)
	for count > 0 {
		batch = append(batch, game.DrawPiece(0, true))
		count--
	}

	game.stateChangeCallback(game.generateDisplay())
	return batch
}

func (game *ChrominoGame) generateDisplay() StateDisplay {
	return StateDisplay{
		DeckSize: len(game.deck.chrominos),
	}
}
