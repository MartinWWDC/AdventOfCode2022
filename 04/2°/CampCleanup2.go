package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//safe := []int{}
	counter := 0
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		//fmt.Println(arr)
		A := strings.Split(arr[0], "-")

		i, _ := strconv.Atoi(A[0])
		f, _ := strconv.Atoi(A[1])
		arr1 := genArr(i, f)
		B := strings.Split(arr[1], "-")
		i, _ = strconv.Atoi(B[0])
		f, _ = strconv.Atoi(B[1])
		arr2 := genArr(i, f)

		for _, i := range arr1 {
			if isIn(arr2, i) {
				/*
					if !isIn(safe, i) {
						safe = append(safe, i)
					}*/
				counter++
				break
			}
		}

	}
	//fmt.Println(safe)
	//fmt.Println(len(safe))
	fmt.Println(counter)
}
func genArr(min int, max int) (arr []int) {
	for i := min; i <= max; i++ {
		arr = append(arr, i)
	}
	return
}
func isIn(arr []int, str int) bool {
	for _, r := range arr {
		if r == str {
			return true
		}
	}
	return false
}
