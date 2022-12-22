package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	findTheStart()
}

func findTheStart() {
	var h int
	var str string

	Scanner := bufio.NewScanner(os.Stdin)
	for Scanner.Scan() {
		str = Scanner.Text()
		h = len(str) / 14
		fmt.Println(h, " ", len(str))
	}
	var arr []string
	counter := 0
	c := false
	for i := 0; i <= 14*h; i++ {
		sub := str[i : i+14]
		for i1, r1 := range sub {
			for i2, r2 := range sub {
				if i1 != i2 {
					if string(r1) == string(r2) {
						c = true
						counter = i + 14
						break

					} else {
						//fmt.Println(string(r1), "!=", string(r2))
					}
				}
			}

		}
		if c == false {
			fmt.Println("found: ", sub)

			fmt.Println(counter + 1)
			break
		}
		c = false

	}

	fmt.Println(arr)
}
