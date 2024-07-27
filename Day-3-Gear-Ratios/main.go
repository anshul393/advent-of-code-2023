package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)

	lines := []string{}

	for scan.Scan() {
		line := scan.Text()
		lines = append(lines, line)
	}

	sm := 0

	mp := make(map[string][]int)

	for row, line := range lines {
		num := 0
		start_col := -1
		end_col := -1
		for col, word := range line {
			if nm, err := strconv.Atoi(string(word)); err == nil {
				num = num*10 + nm
				if start_col == -1 {
					start_col = col
				}
				end_col = col
			} else {
				// check if number gets included
				if IsValid(row, start_col, end_col, lines) {
					// fmt.Println(num)
					sm += num

				}

				for _, st := range IsValidMul(row, start_col, end_col, lines) {
					mp[fmt.Sprintf("row-%d-col-%d", st.Row, st.Col)] = append(mp[fmt.Sprintf("row-%d-col-%d", st.Row, st.Col)], num)
				}

				num = 0
				start_col = -1
				end_col = -1

			}
		}

		if IsValid(row, start_col, end_col, lines) {
			sm += num
			// fmt.Println(num)
		}

		for _, st := range IsValidMul(row, start_col, end_col, lines) {
			mp[fmt.Sprintf("row-%d-col-%d", st.Row, st.Col)] = append(mp[fmt.Sprintf("row-%d-col-%d", st.Row, st.Col)], num)
		}

	}

	smm := 0
	for _, nums := range mp {
		smm += product(nums)
	}

	fmt.Println(sm, smm)
}

func product(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			ans += nums[i] * nums[j]
		}
	}

	return ans
}

func IsValid(row, s_col, e_col int, lines []string) bool {
	flag := false
	// row-1
	if row-1 >= 0 {
		for j := s_col - 1; j <= e_col+1; j++ {
			if j >= 0 && j < len(lines[0]) {
				if string(lines[row-1][j]) != "." && !IsNumeric(string(lines[row-1][j])) {
					flag = true
					return flag
				}
			}
		}
	}
	// row
	if s_col-1 >= 0 && string(lines[row][s_col-1]) != "." && !IsNumeric(string(lines[row][s_col-1])) {
		flag = true
		return flag
	}

	if e_col+1 < len(lines[0]) && string(lines[row][e_col+1]) != "." && !IsNumeric(string(lines[row][e_col+1])) {
		flag = true
		return flag
	}

	// row+1
	if row+1 < len(lines) {
		for j := s_col - 1; j <= e_col+1; j++ {
			if j >= 0 && j < len(lines[0]) {
				if string(lines[row+1][j]) != "." && !IsNumeric(string(lines[row+1][j])) {
					flag = true
					return flag
				}
			}
		}
	}

	return false
}

type star struct {
	Row int
	Col int
}

func IsValidMul(row, s_col, e_col int, lines []string) []star {

	r := make([]star, 0)

	if row-1 >= 0 {
		for j := s_col - 1; j <= e_col+1; j++ {
			if j >= 0 && j < len(lines[0]) {
				if string(lines[row-1][j]) == "*" {
					r = append(r, star{Row: row - 1, Col: j})
				}
			}
		}
	}
	// row
	if s_col-1 >= 0 && string(lines[row][s_col-1]) == "*" {
		r = append(r, star{Row: row, Col: s_col - 1})
	}

	if e_col+1 < len(lines[0]) && string(lines[row][e_col+1]) == "*" {
		r = append(r, star{Row: row, Col: e_col + 1})
	}

	// row+1
	if row+1 < len(lines) {
		for j := s_col - 1; j <= e_col+1; j++ {
			if j >= 0 && j < len(lines[0]) {
				if string(lines[row+1][j]) == "*" {
					r = append(r, star{Row: row + 1, Col: j})
				}
			}
		}
	}

	return r
}

func IsNumeric(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}
