package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://okscoring.com",
	}

		// Channels are typed
		c := make(chan string) // a channel that passes strings



	// for _, link := range links {
	// 	checkLinkBlocking(link)
	// }

	// for _, link := range links {
	// 	// Spin a new thread (routine) for each check
	// 	go checkLinkBlockingWithChannel(link, c)
	// }
	// Even though we called that function multiple times, we still exit at this point before the functions are completed, because they are in a separate routine, and this one (the main routine) has completed at this point
	// So We need to use channels to communicate between routines


	// Pass any messages passed from c into fmt Println
	// this line is blocking, so main routine sleeps until a value is passed
	// the main routine wakes up when a message is received
	// then this code executes, and the program terminates because the blocked code has resolved. The routine only waits for 1 value to come through the channel
	// fmt.Println(<- c)
	// Duplicating ^ will wait for 2 messages etc.

	// In this example, if we set up n + 1 channel message listeners in a row (where n is the length of the array above) the program would hang, because there would not be enough messages coming back in this program

	// So instead going to loop and call it
	// Note, if we set this up in the for loop above, that would block the loop on each iteration

	// for i := 0; i < len(links); i ++ {
	// 	fmt.Println(<- c)
	// }

	for _, link := range links {
		// Spin a new thread (routine) for each check
		go checkLink(link, c)
	}

	// This will loop every time a message comes back through from the channel
	// On each iteration it will create a new routine (thread) and wait for a message to return
	// from the channel, then pass the return back to checkLink to continuously check the link
	for link := range c { // range with a channel basically will keep this loop alive and running for each message from the channel
		// Use a function literal to sleep the new go routine rather than the
		// main routine (because that would block the main routine)
		go func (l string) {
			time.Sleep(time.Second * 2)
			// Running like this captures the link variable in the anon function (function literal)
			// I believe this would basically create a closure?
			// The runtime behavior is that after the initial loop, the value
			// used for link will alwyas be the same as the final value from the initial set of iterations
			// This is not intuitive because go makes a copy when passing values, but in this case
			// it is always referencing the same memory address
			// checkLink(link, c)

			// Oh I get it, because previously we were referencing the variable IN the new go routine code which is a nono because that will share a memory address across the routines
			// By passing to the anon function, you make a copy before it goes into the new routine
			checkLink(l, c)
		}(link) // IIFE, immediately invoked function expression
	}
}

// func checkLinkBlocking(link string) bool {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		fmt.Println(link, "might be down")
// 		return false
// 	}
// 	fmt.Println(link, "is up")
// 	return true
// }

func checkLinkBlockingWithChannel(link string, c chan string) {
	_, err := http.Get(link)

	var message string
	if err != nil {
		message = link + " might be down"
	} else {
		message = link + " is up"
	}
	// Pass this string to the channel
	c <- message
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	var message string
	if err != nil {
		message = link + " might be down"
	} else {
		message = link + " is up"
	}
	fmt.Println(message)
	// Pass this string to the channel
	c <- link
}
