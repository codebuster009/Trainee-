package main

import "fmt"

func main() {
	cards := deck{"Ace of Spades", newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Spades"
}
