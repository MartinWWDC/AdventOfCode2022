package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	counter := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		//fmt.Println(arr)
		A := strings.Split(arr[0], "-")

		iA, _ := strconv.Atoi(A[0])
		fA, _ := strconv.Atoi(A[1])

		B := strings.Split(arr[1], "-")
		iB, _ := strconv.Atoi(B[0])
		fB, _ := strconv.Atoi(B[1])
		fmt.Println("A: ", iA, "  ", fA, "B: ", iB, "  ", fB)

		if (iA <= iB && fB <= fA) || (iB <= iA && fA <= fB) {
			fmt.Println(arr[0], " inside ", arr[1])
			counter++
		}
	}
	fmt.Println(counter)
}
