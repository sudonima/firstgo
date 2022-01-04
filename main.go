package main

func main() {
	sl := newDeck()
	// // sl.chap()
	m, _ := deal(sl, 2)
	// m.chap()
	// fmt.Println("opppsss")
	// n.chap()
	m.SaveToFile()

	Hey := NewDeckFromFile()
	Hey.chap()

	shuf := newDeck()
	shuf.Shuffle()
	shuf.chap()
}

func asgar() string {
	return "NOOOOO"
}
