package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Part struct {
	number int
	gears  []Position
}

func Day3Part1() {
	file, err := os.Open("input/day3.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	sum := 0

	for i, line := range lines {
		splitedLine := strings.Split(line, "")

		previousLine := []string{}
		nextLine := []string{}

		emptyRow := strings.Repeat(".", len(splitedLine))
		emptyRowSplited := strings.Split(emptyRow, "")

		if i == 0 {
			previousLine = emptyRowSplited
		} else {
			previousRowSplited := strings.Split(lines[i-1], "")
			previousLine = previousRowSplited
		}

		if i == len(splitedLine)-1 {
			nextLine = emptyRowSplited
		} else {
			previousRowSplited := strings.Split(lines[i+1], "")
			nextLine = previousRowSplited
		}

		captureNumMode := false
		strNums := map[string]bool{
			"0": true,
			"1": true,
			"2": true,
			"3": true,
			"4": true,
			"5": true,
			"6": true,
			"7": true,
			"8": true,
			"9": true,
		}

		top := []string{}
		mid := []string{}
		bottom := []string{}

		capture := [3][]string{}
		parts := [][3][]string{}

		for j, char := range splitedLine {

			if !captureNumMode && strNums[char] {
				captureNumMode = true

				if j == 0 {
					top = append(top, ".")
					mid = append(mid, ".")
					bottom = append(bottom, ".")
				} else {
					top = append(top, previousLine[j-1])
					mid = append(mid, splitedLine[j-1])
					bottom = append(bottom, nextLine[j-1])
				}
			}

			if !strNums[char] && len(mid) > 1 {
				captureNumMode = false

				top = append(top, previousLine[j])
				mid = append(mid, char)
				bottom = append(bottom, nextLine[j])

				capture[0] = top
				capture[1] = mid
				capture[2] = bottom
				// fmt.Println(capture)
				parts = append(parts, capture)

				top = []string{}
				mid = []string{}
				bottom = []string{}
			}

			if captureNumMode {
				top = append(top, previousLine[j])
				mid = append(mid, char)
				bottom = append(bottom, nextLine[j])

				if j == len(splitedLine)-1 {
					captureNumMode = false
					if strNums[char] {
						top = append(top, ".")
						mid = append(mid, ".")
						bottom = append(bottom, ".")
					}

					capture[0] = top
					capture[1] = mid
					capture[2] = bottom

					// fmt.Println(capture)

					parts = append(parts, capture)
					capture = [3][]string{}

					top = []string{}
					mid = []string{}
					bottom = []string{}
				}
			}
		}

		for _, part := range parts {
			for x := 0; x < len(part); x++ {
				for y := 0; y < len(part[x]); y++ {
					symbols := "@#$%&*-/+="
					if strings.Contains(symbols, part[x][y]) {
						middle := part[1]
						trim := middle[1 : len(middle)-1]
						join := strings.Join(trim, "")
						con, _ := strconv.Atoi(join)
						sum += con
					}
				}
			}
		}
	}

	fmt.Println("sum: ", sum)
}

type PartDetail struct {
	number int
	line   int
	start  int
	end    int
}

func Day3Part2() {

	// testInput := []string{
	// 	"467..114..",
	// 	"...*......",
	// 	"..35..633.",
	// 	"......#...",
	// 	"617*......",
	// 	".....+.58.",
	// 	"..592.....",
	// 	"......755.",
	// 	"...$.*....",
	// 	".664.598..",
	// }

	file, err := os.Open("input/day3.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	strNums := map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, "."+line+".")
	}

	// for _, l := range testInput {
	// 	lines = append(lines, "."+l+".")
	// }

	preparedLines := []string{}

	for i, line := range lines {
		if i == 0 {
			preparedLines = append(preparedLines, strings.Repeat(".", len(line)))
		}
		preparedLines = append(preparedLines, line)
		if i == len(lines)-1 {
			preparedLines = append(preparedLines, strings.Repeat(".", len(line)))
		}
	}

	gearIndexes := [][]int{}

	productSum := 0

	for i, line := range preparedLines {
		// scanning gear indexes
		symbols := []int{}
		for j, char := range line {
			if char == rune('*') {
				symbols = append(symbols, j)
			}
		}
		if len(symbols) == 0 {
			continue
		}

		fmt.Println(symbols)

		gearIndexes = append(gearIndexes, symbols)
		numbersInLine := []PartDetail{}
		numbersAbove := []PartDetail{}
		numbersBelow := []PartDetail{}

		// geting numbers in line
		captureMode := false
		inlineCapture := []string{}
		for j, char := range line {
			if !strNums[char] {
				captureMode = false

				if len(inlineCapture) > 0 {
					join := strings.Join(inlineCapture, "")
					con, _ := strconv.Atoi(join)
					partDetails := PartDetail{number: con, line: i, start: j - len(inlineCapture) - 1, end: j}
					numbersInLine = append(numbersInLine, partDetails)
					inlineCapture = nil
				}

				continue
			}
			if strNums[char] {
				captureMode = true
			}
			if captureMode {
				inlineCapture = append(inlineCapture, string(char))
			}
		}

		if len(symbols) > 0 {
			// checking numbers above
			aboveCaptureMode := false
			aboveCapture := []string{}
			for j, char := range preparedLines[i-1] {
				if !strNums[char] {
					aboveCaptureMode = false

					if len(aboveCapture) > 0 {
						join := strings.Join(aboveCapture, "")
						con, _ := strconv.Atoi(join)
						partDetails := PartDetail{number: con, line: i - 1, start: j - len(aboveCapture) - 1, end: j}
						numbersAbove = append(numbersAbove, partDetails)
						aboveCapture = nil
					}

					continue
				}
				if strNums[char] {
					aboveCaptureMode = true
				}
				if aboveCaptureMode {
					aboveCapture = append(aboveCapture, string(char))
				}
			}

			// checking numbers below
			belowCaptureMode := false
			belowCapture := []string{}
			for j, char := range preparedLines[i+1] {
				if !strNums[char] {
					belowCaptureMode = false

					if len(belowCapture) > 0 {
						join := strings.Join(belowCapture, "")
						con, _ := strconv.Atoi(join)
						partDetails := PartDetail{number: con, line: i - 1, start: j - len(belowCapture) - 1, end: j}
						numbersBelow = append(numbersBelow, partDetails)
						belowCapture = nil
					}

					continue
				}
				if strNums[char] {
					belowCaptureMode = true
				}
				if belowCaptureMode {
					belowCapture = append(belowCapture, string(char))
				}
			}
		}

		// fmt.Println(numbersBelow)

		for _, symbolIndex := range symbols {
			// check symbols collitions vertically
			for _, aboveNumber := range numbersAbove {
				for _, belowNumber := range numbersBelow {
					if aboveNumber.start <= belowNumber.end && aboveNumber.end >= belowNumber.start {
						if aboveNumber.start <= symbolIndex && symbolIndex <= aboveNumber.end &&
							belowNumber.start <= symbolIndex && symbolIndex <= belowNumber.end {
							fmt.Println("v", ",", "line:", i, ", symbol pos:", symbolIndex, ", above:", aboveNumber.number, "-", aboveNumber.start, "-", aboveNumber.end, ", below:", belowNumber.number, "-", belowNumber.start, "-", belowNumber.end)
							productSum += aboveNumber.number * belowNumber.number
						}
					}
				}
			}

			// check symbols collitions horizontally
			for _, v := range numbersInLine {
				for _, w := range numbersInLine {
					if v.end == w.start && v.end == symbolIndex && w.start == symbolIndex {
						fmt.Println("h", ",", "line:", i, ", symbol pos:", symbolIndex, ", first:", v.number, "-", v.start, "-", v.end, ", second:", w.number, "-", w.start, "-", w.end)
						productSum += v.number * w.number
					}
				}
			}

			// check symbols collitions diagonally
			for _, v := range numbersInLine {
				for _, aboveNumber := range numbersAbove {
					if v.end == aboveNumber.start && v.end == symbolIndex && aboveNumber.start == symbolIndex ||
						v.start == aboveNumber.end && v.start == symbolIndex && aboveNumber.end == symbolIndex {
						fmt.Println("d", ",", "line:", i, ", symbol pos:", symbolIndex, ", first:", v.number, "-", v.start, "-", v.end, ", second:", aboveNumber.number, "-", aboveNumber.start, "-", aboveNumber.end)
						productSum += v.number * aboveNumber.number
					}
				}

				for _, belowNumber := range numbersBelow {
					if v.end == belowNumber.start && v.end == symbolIndex && belowNumber.start == symbolIndex ||
						v.start == belowNumber.end && v.start == symbolIndex && belowNumber.end == symbolIndex {
						fmt.Println("d", ",", "line:", i, ", symbol pos:", symbolIndex, ", first:", v.number, "-", v.start, "-", v.end, ", second:", belowNumber.number, "-", belowNumber.start, "-", belowNumber.end)
						productSum += v.number * belowNumber.number
					}
				}
			}

			// check symbols that have both numbers above
			for _, na1 := range numbersAbove {
				for _, na2 := range numbersAbove {
					if na1.start == symbolIndex && na2.end == symbolIndex {
						fmt.Println("t", ",", "line:", i, ", symbol pos:", symbolIndex, ", first:", na1.number, "-", na1.start, "-", na1.end, ", second:", na2.number, "-", na2.start, "-", na2.end)
						productSum += na1.number * na2.number
					}
				}

			}

			// check symbols that have both numbers below
			for _, nb1 := range numbersBelow {
				for _, nb2 := range numbersBelow {
					if nb1.start == symbolIndex && nb2.end == symbolIndex {
						fmt.Println("t", ",", "line:", i, ", symbol pos:", symbolIndex, ", first:", nb1.number, "-", nb1.start, "-", nb1.end, ", second:", nb2.number, "-", nb2.start, "-", nb2.end)
						productSum += nb1.number * nb2.number
					}
				}
			}
		}

		symbols = nil
	}

	fmt.Println(productSum)

}
