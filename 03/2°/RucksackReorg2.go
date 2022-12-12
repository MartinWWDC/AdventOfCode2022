package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	//var arr []rune
	var arrBack []string
	i := 0
	for scanner.Scan() {
		data := scanner.Text()
		arrBack = append(arrBack, data)
		i++
		if i == 3 {
			i = 0
			fmt.Println(arrBack)
			fmt.Println(len(arrBack))
			check := false
			for _, a0 := range arrBack[0] {
				for _, a1 := range arrBack[1] {
					for _, a2 := range arrBack[2] {
						if a0 == a1 && a1 == a2 {

							switch {
							case a1 >= 'a' && a1 <= 'z':
								fmt.Println(string(a1))
								fmt.Println(a1 - 96)
								counter += int(a1 - 96)
								check = true
							case a1 >= 'A' && a1 <= 'Z':
								fmt.Println("ciao")
								fmt.Println(string(a1))

								fmt.Println(a1 - 38)
								counter += int(a1 - 38)
								check = true

							}
						}
						if check {
							break
						}
					}
					if check {
						break
					}
				}
				if check {
					break
				}
			}
			arrBack = nil
		}

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
