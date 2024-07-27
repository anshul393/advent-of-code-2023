package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)

	sum := 0

	for s.Scan() {
		word := s.Text()
		sum += extract1(word)
	}

	fmt.Println(sum)
}

func extract(word string) int {
	num1, num2 := -1, -1

	for _, char := range word {
		if num, err := strconv.Atoi(string(char)); err == nil {
			if num1 == -1 {
				num1 = num
			} else {
				num2 = num
			}
		}
	}

	ans := 0

	if num1 == -1 {
		ans = 0
	} else if num2 != -1 {
		ans = 10*num1 + num2
	} else {
		ans = 10*num1 + num1
	}

	return ans
}

func extract1(word string) int {
	mp := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	num1, num2 := -1, -1

	l := len(word)

	for idx, char := range word {
		ans := -1
		if num, err := strconv.Atoi(string(char)); err == nil {
			ans = num
		} else {
			if val, ok := mp[word[idx:min(idx+3, l)]]; ok {
				ans = val
			} else if val, ok := mp[word[idx:min(idx+4, l)]]; ok {
				ans = val
			} else if val, ok := mp[word[idx:min(idx+5, l)]]; ok {
				ans = val
			}
		}

		if ans != -1 {
			if num1 == -1 {
				num1 = ans
			} else {
				num2 = ans
			}
		}
	}

	ans := 0

	if num1 == -1 {
		ans = 0
	} else if num2 != -1 {
		ans = 10*num1 + num2
	} else {
		ans = 10*num1 + num1
	}

	return ans
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}
