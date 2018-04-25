package main

import "math/rand"

type deck struct {
	// if true, the card is drawn.
	cards [52]bool
}

func newDeck() *deck {
	return &deck{}
}

func (d *deck) draw() int {
	var drawnN int
	for {
		n := rand.Intn(52)
		if d.cards[n] {
			continue
		} else {
			d.cards[n] = true
			drawnN = n
			break
		}
	}
	return drawnN
}
