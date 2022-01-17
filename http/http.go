package main

import (
	"fmt"
)

type logWriter struct{}

func main() {
	// resp, err := http.Get("https://kassidyrhodesphotography.com")

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// responseBytes := make([]byte, 99999) // Create a byte slice of size 99,999
	// // Read function reads data into a slice until it is full
	// // So maybe this could be useful for chunking?

	// // Pass the byte slice to the Read function to read data into it.
	// resp.Body.Read(responseBytes)

	// fmt.Println(string(responseBytes))

	// // This basically does the same thing
	// // it writes the Readable Body to Stdout
	// io.Copy(os.Stdout, resp.Body)

	fmt.Println(test(10))


}


// Declaring the return type with a name creates that variable in this scope with a default value
func test(i int) (result int) {
	return i + result
}

// logWriter now implements the Writer interface
// This is interesting, this could easily be used as a middleware solution for logging etc
// Although I am not sure if you can only read once from a source
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
