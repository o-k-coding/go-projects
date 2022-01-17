package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpMethod := "GET"
	url := "https://api.github.com"

	client := http.Client{}

	// response, err := client.Get(url)
	// or
	request, err := http.NewRequest(httpMethod, url, nil)

	request.Header.Set("Accept", "application/xml")

	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}


	//  A defer statement defers the execution of a function until the surrounding function returns.
	// Note: any arguments passed to this call are immediately evaluated, but the call is not made until the function has completed
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
