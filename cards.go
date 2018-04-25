package main

import "strconv"

type cards []card

func (c cards) reveal() (sumExceptAce, countAce int) {
	for _, card := range c {
		if card.strength() == 1 {
			countAce++
		} else {
			sumExceptAce += card.strength()
		}
	}
	return sumExceptAce, countAce
}

func (c cards) strength() int {
	var sum = 0
	sumExceptAce, countAce := c.reveal()
	switch countAce {
	case 0:
		sum = sumExceptAce
	case 1:
		if sumExceptAce <= 10 {
			sum = sumExceptAce + 11
		} else {
			sum = sumExceptAce + 1
		}
	default:
		if sumExceptAce+countAce > 11 {
			sum = sumExceptAce + countAce
		} else {
			sum = sumExceptAce + countAce + 10
		}
	}
	if sum > 21 {
		return 0
	}
	return sum
}

func (c cards) toString() string {
	var s string
	for _, card := range c {
		s += card.toString()
	}
	return s
}

func newGame() (c cards) {
	c = make([]card, 0, 12)
	for i := 0; i < 2; i++ {
		c = append(c, card{id: usedDeck.draw()})
	}
	return c
}

type card struct {
	id int
}

func (c card) suit() string {
	switch c.id % 4 {
	case 0:
		return "C"
	case 1:
		return "D"
	case 2:
		return "H"
	default:
		return "S"
	}
}

// Aの場合1を返す
func (c card) strength() int {
	if c.id >= 40 {
		return 10
	} else {
		return c.id/4 + 1
	}
}

func (c card) number() string {
	switch c.id / 4 {
	case 0:
		return "A"
	case 9:
		return "T"
	case 10:
		return "J"
	case 11:
		return "Q"
	case 12:
		return "K"
	default:
		return strconv.Itoa(c.id/4 + 1)
	}
}

func (c card) toString() string {
	var s string
	s += c.suit()
	s += c.number()
	return s
}
