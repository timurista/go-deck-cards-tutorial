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
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "King", "Queen",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, "	", card)
	}
}

func (d deck) toString() string {
	deckString := []string(d)

	// for i, card := range d {
	// 	fmt.Println(i+1, "	", card)

	// }
	return strings.Join(deckString, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		println("Error: " + err.Error())
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	for i := range d {
		newPostion := r.Intn(len(d) - 1)
		d[i], d[newPostion] = d[newPostion], d[i]
	}
}
