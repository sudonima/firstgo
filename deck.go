package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join(([]string)(d), ",")
}

func (d deck) SaveToFile() error {
	return ioutil.WriteFile("nima", []byte(d.toString()), 0777)
}

func NewDeckFromFile() deck {
	bs, err := ioutil.ReadFile("nima")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixMilli())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
