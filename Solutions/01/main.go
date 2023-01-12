package main

import "fmt"

type Card struct {
	suit  string
	value string
}

var suits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
var values = []string{"Ace", "Two", "Three", "Four"}

func main() {
	card, err := newCard("Two", "Hearts")

	if err != nil {
		fmt.Printf("%v", err)
	}

	card.print()
}

func newCard(v string, s string) (Card, error) {

	//Flag variables for error catching
	var invalidValue = true
	var invalidSuit = true
	var card Card

	for _, val := range values {
		if val == v {
			invalidValue = false //Sets error flag to be false (no error) if val equals v
			break
		}
	}

	if invalidValue {
		return card, fmt.Errorf("Invalid Value was past in as an argument")
	}

	for _, suit := range suits {
		if suit == s {
			invalidSuit = false //Sets error flag to be false (no error) if suit equals s
			break
		}
	}

	if invalidSuit {
		return card, fmt.Errorf("Invalid Suit was past in as an argument")
	}

	card.value = v
	card.suit = s

	return card, nil
}

func (c Card) print() {
	fmt.Println(c.value + " of " + c.suit)
}
