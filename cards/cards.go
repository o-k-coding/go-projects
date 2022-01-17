package main

func main() {
	// Slice can only have one type for all values
	cards := newDeck()
	cards.saveToFile("cards.csv")

	readCards := readDeckFromFile("cards.csv")
	readCards.print()

	readCards.shuffle()
	readCards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
