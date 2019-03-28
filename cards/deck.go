package main

import "fmt"

// Create a new type of "deck"
// which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d { // 1. d is similar to this 2. by convention, use single letter of the type
		fmt.Println(i, card)
	}
}
