package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func calcCalibrationValue(values []int) int {
	if len(values) > 1 {
		numsJoin := strconv.Itoa(values[0]) + strconv.Itoa(values[len(values)-1])
		result, err := strconv.ParseInt(numsJoin, 10, 32)
		if err != nil {
			return 0
		}
		return int(result)
	} else {
		numsJoin := strconv.Itoa(values[0]) + strconv.Itoa(values[0])
		result, err := strconv.ParseInt(numsJoin, 10, 32)
		if err != nil {
			return 0
		}
		return int(result)
	}
}

func GetCalibrationValue(line string) int {
	lineArr := strings.Split(line, "")
	numsArr := []int{}

	for i := 0; i < len(lineArr); i++ {
		int64value, err := strconv.ParseInt(lineArr[i], 10, 32)
		if err != nil {
			continue
		}
		num := int(int64value)
		numsArr = append(numsArr, num)
	}

	return calcCalibrationValue(numsArr)
}

func GetCalibrationValueWithChars(line string) int {
	matches := []string{}

	re := regexp2.MustCompile(`(?=(one|two|three|four|five|six|seven|eight|nine|\d{1}))`, 0)
	for m, _ := re.FindStringMatch(line); m != nil; m, _ = re.FindNextMatch(m) {
		matches = append(matches, m.Groups()[1].Capture.String())
	}

	numsArr := []int{}
	for i := 0; i < len(matches); i++ {
		switch matches[i] {
		case "one":
			numsArr = append(numsArr, 1)
			break
		case "two":
			numsArr = append(numsArr, 2)
			break
		case "three":
			numsArr = append(numsArr, 3)
			break
		case "four":
			numsArr = append(numsArr, 4)
			break
		case "five":
			numsArr = append(numsArr, 5)
			break
		case "six":
			numsArr = append(numsArr, 6)
			break
		case "seven":
			numsArr = append(numsArr, 7)
			break
		case "eight":
			numsArr = append(numsArr, 8)
			break
		case "nine":
			numsArr = append(numsArr, 9)
			break
		default:
			numparsed, err := strconv.Atoi(matches[i])
			if err != nil {
				return 0
			}
			n := int(numparsed)
			numsArr = append(numsArr, n)
		}
	}

	return calcCalibrationValue(numsArr)
}

func calcSum(values []int) int {
	result := 0

	for i := 0; i < len(values); i++ {
		result += values[i]
	}

	return result
}

func Day1Part1() {
	file, err := os.Open("input/day1.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, GetCalibrationValue(line))
	}

	result := calcSum(results)
	fmt.Printf("Result: %d\n", result)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func Day1Part2() {
	file, err := os.Open("input/day1.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		value := GetCalibrationValueWithChars(line)
		results = append(results, value)
	}

	result := calcSum(results)
	fmt.Printf("Result: %d\n", result)
}
