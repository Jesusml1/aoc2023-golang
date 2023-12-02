package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id   int
	bags []Bag
}

func Day2Part1() {
	file, err := os.Open("input/day2.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := []Game{}

	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, ParseGame(line))
	}

	idSumValidGames := 0

	for _, game := range games {
		PrintGame(game)
		if CheckValidGame(game) {
			idSumValidGames += game.id
		}
	}
	fmt.Println(idSumValidGames)
}

func Day2Part2() {
	file, err := os.Open("input/day2.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := []Game{}

	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, ParseGame(line))
	}

	powerSum := 0

	for _, game := range games {
		powerSum = powerSum + CalcPowerSet(game)
	}

	fmt.Println(powerSum)
}

func ParseGame(line string) Game {
	mid := strings.Split(line, ":")
	gameNumber := strings.Split(mid[0], " ")[1]
	bags := strings.Split(mid[1], ";")
	parsedBags := []Bag{}
	for _, bag := range bags {
		red := 0
		green := 0
		blue := 0

		trimedBag := strings.Trim(bag, " ")
		cubes := strings.Split(trimedBag, ", ")
		for i := 0; i < len(cubes); i++ {
			pair := cubes[i]
			splitPair := strings.Split(pair, " ")
			cubeNum, _ := strconv.Atoi(splitPair[0])
			switch splitPair[1] {
			case "red":
				red = cubeNum
				break
			case "green":
				green = cubeNum
				break
			case "blue":
				blue = cubeNum
				break
			}
		}

		parsedBags = append(parsedBags, Bag{red: red, green: green, blue: blue})
	}

	gameId, _ := strconv.Atoi(gameNumber)

	return Game{id: gameId, bags: parsedBags}
}

func PrintGame(game Game) {
	fmt.Println("game id:", game.id)
	for _, bag := range game.bags {
		fmt.Println(bag)
	}
}

func CheckValidGame(game Game) bool {
	for _, bag := range game.bags {
		if bag.red > 12 || bag.green > 13 || bag.blue > 14 {
			return false
		}
	}
	return true
}

func CalcPowerSet(game Game) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, bag := range game.bags {
		if bag.red > maxRed {
			maxRed = bag.red
		}
		if bag.green > maxGreen {
			maxGreen = bag.green
		}
		if bag.blue > maxBlue {
			maxBlue = bag.blue
		}
	}

	return maxRed * maxGreen * maxBlue
}
