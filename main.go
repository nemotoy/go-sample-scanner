package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	ch := make(chan string)

	go func() {
		for {
			select {
			case str := <-ch:
				fmt.Println("result: ", str)
			}
		}
	}()

	for scanner.Scan() {
		ch <- scanner.Text()
	}
}
