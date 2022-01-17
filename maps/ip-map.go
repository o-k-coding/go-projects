package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readIPsFromFile(fileName string) map[string]int {
	ipBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading IPs file: ", err)
		os.Exit(1)
	}

	ipMap := make(map[string]int)
	ip := ""

	for i, char := range ipBytes {
		str := string(char)
		if str == "\n" {
			ipMap[ip] = i
			fmt.Println("Adding ip", ip)
			ip = ""
		} else {
			ip += str
		}
	}

	fmt.Println("Loaded", len(ipMap), "ip addresses")
	return ipMap
}
