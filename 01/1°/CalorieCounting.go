package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type elf struct {
	id     int
	amount int
}

func main() {
	readData()
}
func readData() {
	Scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	i := 1
	var max elf
	for Scanner.Scan() {
		if Scanner.Text() != "" {
			tmp, _ := strconv.Atoi(Scanner.Text())

			counter += tmp
		} else {
			if i == 1 {
				max = elf{i, counter}
			} else {
				if counter > max.amount {
					max = elf{i, counter}
				}
			}
			counter = 0

			i++
		}

	}
	if counter > max.amount {
		max = elf{i, counter}
	}
	fmt.Println(max)

}
