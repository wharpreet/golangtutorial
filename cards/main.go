package main

func main() {
	cards := newDeck()
	cards.shuffle()
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_card")
	cards.print()

	//hand, remaingCards := deal(cards, 10)
	//hand.print()
	//remaingCards.print()
}
