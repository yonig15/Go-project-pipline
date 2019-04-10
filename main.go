package main

func main() { // Has to be first matching package main
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
