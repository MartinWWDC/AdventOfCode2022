package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	var arr []rune
	for scanner.Scan() {
		data := scanner.Text()
		sec0 := data[0 : len(data)/2]
		sec1 := data[len(data)/2:]
		for s := range sec0 {
			for d := range sec1 {
				if !contain(rune(sec0[s]), arr) {
					if sec0[s] == sec1[d] {
						switch {
						case sec0[s] >= 'a' && sec0[s] <= 'z':
							fmt.Println(string(sec0[s]))
							fmt.Println(sec0[s] - 96)
							counter += int(sec0[s] - 96)
							arr = append(arr, rune(sec0[s]))
						case sec0[s] >= 'A' && sec0[s] <= 'Z':
							fmt.Println("ciao")
							fmt.Println(string(sec0[s]))

							fmt.Println(sec0[s] - 38)
							counter += int(sec0[s] - 38)
							arr = append(arr, rune(sec0[s]))

						}

					}

				}

			}
			fmt.Println(sec0)
			fmt.Println(sec1)
		}
		arr = nil

	}
	fmt.Println(counter)

}
func contain(run rune, arr []rune) bool {
	for _, g := range arr {
		if g == run {
			return true
		}
	}
	return false
}
