package main

import (
	"fmt"
	"strconv"
)

func main() {

	//////// Variable defining
	// #1
	var card0 = "0"

	// #2
	var card1 string
	card1 = "1"

	// #3
	card2 := "2"

	fmt.Println(card0, "-", card1, "-", card2)

	//////// Function call
	card := newCard()
	fmt.Println(card)

	//////// For loop
	cards := []string{"13", "94", "5"}
	cards = append(cards, "22")

	for i, c := range cards {
		fmt.Println("Card"+strconv.Itoa(i)+":", c)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
