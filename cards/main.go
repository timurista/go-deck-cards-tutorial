package main

func main() {
	// cards := newDeck()
	// err := cards.saveToFile("test-deck.deck")
	// if err != nil {
	// 	panic(err)
	// }

	newCards := newDeckFromFile("test-deck.deck")
	newCards.shuffle()
	newCards.print()

	// hand, _ := deal(cards, 5)
	// hand.print()
	// print(cards.toString())

	// remainingCards.print()
	// cards = append(cards, "Five of Spades")
	// fmt.Printf("%+q\n", cards)
}
