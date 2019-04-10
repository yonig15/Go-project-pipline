package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// Which is a slice of strings
type deck []string // custom type

func newDeck() deck { // Create and return list of cards
	cards := deck{}

	cardSuits := []string{"Spades", "Dimonds", "Hearts", "Clubs"} // Slice of card suits
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}

	// Two four loops to itterate between suites and values
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+"of"+" "+suit)
		}
	}
	return cards
}

func (d deck) print() { // Reciver (d deck) d in javascript would be like this
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // Deal the cards
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { // Turn our deck into a string
	return strings.Join([]string(d), ",") // Type conversions
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // Make string using function above and then use type conversion to make bytes
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // bs byte slice
	if err != nil {                      // Error handling
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") // Type conversion and turn and seperate the string back out
	return s
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Generating a new seed for randIntn so it does not use same seed
	r := rand.New(source)

	for i := range d {
		np := r.Intn(len(d) - 1) // Random num between zero and slice

		d[i], d[np] = d[np], d[i] // Take new position and assign to i and take whatever is at i and assign to np
	}

}

//Reciver sets up methods on variables we create
