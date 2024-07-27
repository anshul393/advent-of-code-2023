package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var threshold = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)

	// 12 red cubes, 13 green cubes, and 14 blue cubes
	sum := 0
	power := 0
	for scan.Scan() {
		text := scan.Text()
		game_info := strings.Split(text, ":")
		id, _ := strconv.Atoi(strings.Split(game_info[0], " ")[1])
		if isValid(strings.TrimSpace(game_info[1])) {
			sum += id
		}

		power += getPower(strings.TrimSpace(game_info[1]))
	}

	fmt.Println("sum: ", sum)
	fmt.Println("power: ", power)
}

func isValid(game_info string) bool {
	plays := strings.Split(game_info, ";")
	for _, play := range plays {
		cubes := strings.Split(strings.TrimSpace(play), ",")
		for _, cube := range cubes {
			colors := strings.Split(strings.TrimSpace(cube), " ")
			val, _ := strconv.Atoi(colors[0])
			color := colors[1]

			if val > threshold[color] {
				return false
			}
		}

	}

	return true
}

func getPower(game_info string) int {
	plays := strings.Split(game_info, ";")

	mp := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, play := range plays {
		cubes := strings.Split(strings.TrimSpace(play), ",")
		for _, cube := range cubes {
			colors := strings.Split(strings.TrimSpace(cube), " ")
			val, _ := strconv.Atoi(colors[0])
			color := colors[1]

			if val > mp[color] {
				mp[color] = val
			}
		}

	}

	return mp["green"] * mp["red"] * mp["blue"]
}
