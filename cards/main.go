package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")
}
