package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Vending machine initialized")

	for {
		buf := bufio.NewReader(os.Stdin)
		input, err := buf.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(input))
		}
	}
}
