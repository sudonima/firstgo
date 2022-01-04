package main

import "fmt"

func main() {
	sl := newDeck()
	// sl.chap()
	m, n := deal(sl, 2)
	m.chap()
	fmt.Println("opppsss")
	n.chap()
}

func asgar() string {
	return "NOOOOO"
}
