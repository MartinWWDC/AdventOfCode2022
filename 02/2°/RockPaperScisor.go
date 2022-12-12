package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0

	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		switch arr[1] {
		case "X":
			counter += 0
		case "Y":
			counter += 3
		case "Z":
			counter += 6
		}
		switch arr[0] {
		case "A":
			switch arr[1] {
			case "X":
				counter += 3
			case "Y":
				counter += 1
			case "Z":
				counter += 2
			}
		case "B":
			switch arr[1] {
			case "X":
				counter += 1
			case "Y":
				counter += 2
			case "Z":
				counter += 3
			}
		case "C":
			switch arr[1] {
			case "X":
				counter += 2
			case "Y":
				counter += 3
			case "Z":
				counter += 1
			}
		default:

		}
	}
	fmt.Println(counter)
}
