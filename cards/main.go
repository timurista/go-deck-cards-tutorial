package main

func main() {
	cards := newDeck()
	filename := "test-deck.deck"
	err := cards.saveToFile(filename)
	if err != nil {
		panic(err)
	}

	newCards := newDeckFromFile(filename)
	newCards.shuffle()
	newCards.print()

	// hand, _ := deal(cards, 5)
	// hand.print()
	// print(cards.toString())

	// remainingCards.print()
	// cards = append(cards, "Five of Spades")
	// fmt.Printf("%+q\n", cards)
}
