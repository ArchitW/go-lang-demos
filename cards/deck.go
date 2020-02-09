package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
//which is a slice of string

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack",
		"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// (d deck) receiver of type deck ([]string)
//any variable of type deck now get access to the print method.
//d ? instance of deck variable
// by convention 1-2 letter abbriv. of the type
func (d deck) print() {
	for _, card := range d {
		println(card)
	}
}

//(deck deck) return 2 values both will be of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 777)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		println("Error :", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle()  {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
	newPosition := r.Intn(len(d) -1)
	d[i], d[newPosition] = d[newPosition], d[i]

	}
}
/* reciver vs argument


 */
