package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScrashCard struct {
	id      int
	winNums []int
	nums    []int
}

func Day4Part1() {

	file, err := os.Open("input/day4.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scrashCards := []ScrashCard{}

	for scanner.Scan() {
		line := scanner.Text()
		scrashCards = append(scrashCards, ParseScrashCard(line))
	}

	totalPoints := 0

	for _, card := range scrashCards {
		totalPoints += CalcScrashCardPoints(card)
	}

	fmt.Println("total:", totalPoints)
}

func Day4Part2() {

	file, err := os.Open("input/day4.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scrashCards := []ScrashCard{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		scrashCards = append(scrashCards, ParseScrashCard(line))
	}

	totalCards := processScrashCards(scrashCards, scrashCards, 0)
	fmt.Println("total: ", totalCards)
}

func ParseScrashCard(line string) ScrashCard {

	re := regexp.MustCompile(`\d+`)

	split := strings.Split(line, ":")

	cardId := re.FindAllString(split[0], 1)[0]
	cardIdInt, _ := strconv.Atoi(cardId)

	allNums := strings.Split(split[1], "|")
	winNums := re.FindAllString(allNums[0], -1)
	nums := re.FindAllString(allNums[1], -1)

	winNumsInt := []int{}
	numsInt := []int{}

	for _, n := range winNums {
		parsedInt, e := strconv.Atoi(n)
		if e == nil {
			winNumsInt = append(winNumsInt, parsedInt)
		}
	}

	for _, n := range nums {
		parsedInt, e := strconv.Atoi(n)
		if e == nil {
			numsInt = append(numsInt, parsedInt)
		}
	}

	scrashCard := ScrashCard{id: cardIdInt, winNums: winNumsInt, nums: numsInt}

	return scrashCard
}

func CalcScrashCardPoints(scrashCard ScrashCard) int {
	totalPoints := 0

	for _, winNum := range scrashCard.winNums {
		for _, num := range scrashCard.nums {
			if winNum == num {
				if totalPoints == 0 {
					totalPoints += 1
				} else {
					totalPoints *= 2
				}
			}
		}
	}

	return totalPoints
}

func CalcScrashCardMatches(scrashCard ScrashCard) int {
	totalMatches := 0

	for _, winNum := range scrashCard.winNums {
		for _, num := range scrashCard.nums {
			if winNum == num {
				totalMatches += 1
			}
		}
	}

	return totalMatches
}

func processScrashCards(scrashCards []ScrashCard, mainDeck []ScrashCard, index int) int {
	cardCount := 0

	for _, card := range scrashCards {
		matches := CalcScrashCardMatches(card)
		// fmt.Println("card:", card.id, "matches:", matches, "index:", index)
		if matches > 0 {
			copies := []ScrashCard{}
			cardsIn := 0
			for cardsIn < matches {
				copies = append(copies, mainDeck[card.id+cardsIn])
				cardsIn++
			}
			// fmt.Println("copies:", copies)
			cardCount += processScrashCards(copies, mainDeck, index+1)
		}
		cardCount++
	}

	return cardCount
}
