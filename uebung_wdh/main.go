package main

import "fmt"

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}

// Rekursive Funtion
type deck []string

func newDeck() deck {

	cards := deck{}
}

func (d deck) print() {

	for _, i := range d {
		fmt.Println(i)
	}

}
