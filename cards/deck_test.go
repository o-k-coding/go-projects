package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Error("Expected deck length of 52, but got", len(deck))
	}

	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cardCounts := make(map[string]int)

	for _, card := range deck {
		for _, suit := range suits {
			_, exist := cardCounts[suit]
			if !exist {
				cardCounts[suit] = 0
			}
			if strings.Contains(card, suit) {
					cardCounts[suit] += 1
			}
		}

		for _, value := range values {
			_, exist := cardCounts[value]
			if !exist {
				cardCounts[value] = 0
			}
			if strings.Contains(card, value) {
					cardCounts[value] += 1
			}
		}
	}

	for _, suit := range suits {
		suitCount := cardCounts[suit]
			if suitCount != 13 {
				t.Error("Expected 13 of suit" ,suit, "but got", suitCount)
			}
		}

		for _, value := range values {
			valueCount := cardCounts[value]
			if valueCount != 4 {
				t.Error("Expected 4 of value",value, "but got", valueCount)
			}
		}

		// TODO could I use a radix tree for this instead since they will all have the same possible set of begginning strings?
}

func TestSaveToDeckAndReadDeckFromFile(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)

	deck := newDeck()

	deck.saveToFile(testFile)

	loadedDeck := readDeckFromFile(testFile)

	if len(deck) != len(loadedDeck) {
		t.Error("Expected length of loaded deck to equal that of the original deck", "Original: ", len(deck), "Loaded: ", len(loadedDeck))
	}
	os.Remove(testFile)

}

func TestCreateEvenOddString(t *testing.T) {
	nums := []string{"0 is even", "1 is odd", "2 is even", "3 is odd", "4 is even", "5 is odd", "6 is even", "7 is odd", "8 is even", "9 is odd", "10 is even"}
	expectedResult := strings.Join(nums, "\n")
	result := createEvenOddString(10)
	if result != expectedResult {
		t.Error("Expected string to be", expectedResult, "Bbut got", result)
	}
}
