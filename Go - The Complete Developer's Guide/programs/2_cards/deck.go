package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck' based on a slice of strings (similar to "extending")
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// this is a receiver function
// any variable of type deck now gets access to this function
// in this case, d is the instance that the function is called on
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") // we can convert a deck to []string because that was how deck was originally defined, as a []string
}

func (d deck) saveToFile(filename string) {
	os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bytes), ",")
	return deck(s)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func (d deck) shuffle() {
// 	rand.Seed(time.Now().UnixNano())

// 	for i := range d {
// 		newPos := rand.Intn(len(d) - 1)
// 		d[i], d[newPos] = d[newPos], d[i]
// 	}
// }

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}
