package gamemodel

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

type chrominoDeck struct {
	chrominos []*ChrominoPiece
}

func newDeck() *chrominoDeck {
	newDeck := chrominoDeck{chrominos: make([]*ChrominoPiece, 0)}
	newDeck.init()
	return &newDeck
}

func (deck *chrominoDeck) init() {
	for _, middle := range ChrominoColors {
		for leftIndex, left := range ChrominoColors {
			for _, right := range ChrominoColors[leftIndex:] {
				deck.chrominos = append(deck.chrominos, &ChrominoPiece{
					Colors:   [3]ChrominoColor{left, middle, right},
					Rotation: 0,
					X:        0,
					Y:        0,
				})
				//if c0 != c2 || (c0 == c2 && m == 0) {
				//	deck.chrominos = append(deck.chrominos, ChrominoPiece{
				//		Colors:   [3]ChrominoColor{c0, c1, c2},
				//		Rotation: 0,
				//		X:        0,
				//		Y:        0,
				//	})
				//}
			}
		}
	}
	for _, c := range deck.chrominos {
		fmt.Printf("%+v\n", c)
	}
}

func (deck *chrominoDeck) draw() (chrominoPiece *ChrominoPiece, remaining int) {
	deckSize := len(deck.chrominos)
	if deckSize == 0 {
		return nil, 0
	}
	randIndex := randomdata.Number(deckSize)
	chrominoPiece = deck.chrominos[randIndex]
	deck.chrominos = append(
		deck.chrominos[:randIndex],
		deck.chrominos[randIndex+1:]...)
	remaining = deckSize - 1
	return
}
