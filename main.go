package main

import "fmt"

func main() {
	s := "Hello world"
	for _, v := range s {
		fmt.Print(string(v))
	}

}
