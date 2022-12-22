package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	solve(bufio.NewScanner(os.Stdin))
}
func solve(scanner *bufio.Scanner) {
	crates := map[int][]rune{}
	for scanner.Scan() {
		row_raw := scanner.Text()
		if row_raw == "" {
			continue
		}
		row := []rune(row_raw)
		if strings.Contains(row_raw, string('[')) {
			for i := 0; i < len(row); i += 4 {
				if row[i+1] != ' ' {
					crates[i/4+1] = append([]rune{row[i+1]}, crates[i/4+1]...)
				}
			}
		}
		if row[0] == 'm' {
			tokens := strings.Fields(row_raw)
			n_move := str_to_int(tokens[1])
			from := str_to_int(tokens[3])
			to := str_to_int(tokens[5])

			n_from := len(crates[from])
			crates[to] = append(crates[to], crates[from][n_from-n_move:]...)
			crates[from] = crates[from][:n_from-n_move]
		}
	}
	top_crates := ""
	n_pos := len(crates)
	for i := 1; i <= n_pos; i++ {
		n_v := len(crates[i])
		if n_v == 0 {
			continue
		}
		top_crates += string(crates[i][n_v-1])
	}
	fmt.Println(top_crates)
}
