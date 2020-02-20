package main

func main() {
	// var card string = "ace of spades"
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

}

func newCard() string {
	return "🂡"
}
