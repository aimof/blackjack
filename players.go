package main

type dealer struct {
	cards cards
}

func (d *dealer) play() {
	for d.cards.strength() < 17 && d.cards.strength() > 0 {
		d.cards = append(d.cards, card{id: usedDeck.draw()})
	}
}

type player struct {
	cards cards
}

func (p *player) play() {
	p.cards = append(p.cards, card{id: usedDeck.draw()})
}
