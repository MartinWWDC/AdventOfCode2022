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
	//var arr []elf
	i := 1
	max1, max2, max3 := elf{0, 0}, elf{0, 0}, elf{0, 0}
	for Scanner.Scan() {
		if Scanner.Text() != "" {
			tmp, _ := strconv.Atoi(Scanner.Text())
			//fmt.Println(tmp)
			counter += tmp
		} else {
			if counter > max1.amount {
				if counter > max2.amount {
					if counter > max3.amount {
						max1 = max2
						max2 = max3
						max3 = elf{i, counter}

					} else {
						max1 = max2
						max2 = elf{i, counter}

					}
				} else {
					max1 = elf{i, counter}
				}
			}
			//fmt.Println(counter)
			//arr = append(arr, elf{i, counter})
			counter = 0

			i++
		}

	}
	if counter > max1.amount {
		if counter > max2.amount {
			if counter > max3.amount {
				max1 = max2
				max2 = max3
				max3 = elf{i, counter}

			} else {
				max1 = max2
				max2 = elf{i, counter}

			}
		} else {
			max1 = elf{i, counter}
		}

		//arr = append(arr, elf{i, counter})
		//fmt.Println(arr)

	}
	fmt.Println(max1)
	fmt.Println(max2)
	fmt.Println(max3)
	fmt.Println(max1.amount + max2.amount + max3.amount)

}
