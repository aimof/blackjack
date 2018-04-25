package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	usedDeck *deck
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		dealer = dealer{cards: make(cards, 0, 12)}
		player = player{cards: make(cards, 0, 12)}
	)
game:
	for {
		usedDeck = newDeck()
		dealer.cards = newGame()
		player.cards = newGame()
		fmt.Printf("Your cards are %s\n", player.cards.toString())
		fmt.Printf("dealers first card is %s\n", dealer.cards[0].toString())
		fmt.Println("Are you draw? Y/n")
	turn:
		for {
			var in string
			fmt.Scanln(&in)
			switch in {
			case "Y":
				player.play()
				if player.cards.strength() == 0 {
					fmt.Println("You lose!")
					if isContinue() {
						continue game
					} else {
						return
					}
				} else {
					fmt.Printf("Your cards are %s\n", player.cards.toString())
					fmt.Println("Are you draw? Y/n")
					continue
				}
			case "n":
				break turn
			default:
				fmt.Println("input is invalid")
				fmt.Println("Are you draw? Y/n")
				continue
			}
		}
		dealer.play()
		fmt.Printf("Your cards are %s, strength is %d\n", player.cards.toString(), player.cards.strength())
		fmt.Printf("Dealer's cards are %s, strength is %d\n", dealer.cards.toString(), dealer.cards.strength())
		if player.cards.strength() < dealer.cards.strength() {
			fmt.Println("You Lost!")
		} else if player.cards.strength() == dealer.cards.strength() {
			fmt.Println("Draw!")
		} else {
			fmt.Println("You Win!")
		}
		if !isContinue() {
			return
		}
	}
}

func isContinue() bool {
	for {
		fmt.Println("Play again? Y/n")
		var in string
		fmt.Scanln(&in)
		switch in {
		case "Y":
			return true
		case "n":
			return false
		default:
			fmt.Println("input is invalid")
			continue
		}
	}
}
