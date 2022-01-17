package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type deck []string

// Receiver of type deck called d, this is essentially adding a function to the deck type.
// In this case the variable passed into the function is a COPY of the original object. Go functions pass by value unless specifically referenced with &
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// return the hand as the first item, then the rest of the deck as the second item
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readDeckFromFile(fileName string) deck {
	cardBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	cards := strings.Split(string(cardBytes), ",")

	return deck(cards)
}

func (d deck) shuffle() {
	// The seed needs to be different each time, otherwise the randomness will always be the same
	// Funny that they don't have a built in way for this?
	// Could also hash the deck and convert to a number or something. kinda an interesting idea
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	for i := range d {
		newPosition := randGen.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}


func createEvenOddString(n int) string {
	nums := make([]string, n + 1)
	for i := 0; i <= n; i++ {
		iString := strconv.Itoa(i)
		if i % 2 == 0 {
			nums[i] = iString + " is even"
		} else {
			nums[i] = iString + " is odd"
		}
	}

	return strings.Join(nums, "\n")
}
