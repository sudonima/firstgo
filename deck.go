package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	shekl := []string{"ghalb", "gish", "khach"}
	adad := []string{"1", "2", "3", "4"}
	for _, sh := range shekl {
		for _, ad := range adad {
			cards = append(cards, sh+ad)
		}
	}
	return cards
}

func (d deck) chap() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
