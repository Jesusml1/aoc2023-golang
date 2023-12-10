package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	dest int
	src  int
	len  int
}

func Day5Part1() {
	// input := `seeds: 79 14 55 13

	// 	seed-to-soil map:
	// 	50 98 2
	// 	52 50 48

	// 	soil-to-fertilizer map:
	// 	0 15 37
	// 	37 52 2
	// 	39 0 15

	// 	fertilizer-to-water map:
	// 	49 53 8
	// 	0 11 42
	// 	42 0 7
	// 	57 7 4

	// 	water-to-light map:
	// 	88 18 7
	// 	18 25 70

	// 	light-to-temperature map:
	// 	45 77 23
	// 	81 45 19
	// 	68 64 13

	// 	temperature-to-humidity map:
	// 	0 69 1
	// 	1 0 69

	// 	humidity-to-location map:
	// 	60 56 37
	// 	56 93 4`

	data, err := os.ReadFile("input/day5.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	input := string(data)

	lines := strings.Split(input, "\n\n")
	re := regexp.MustCompile(`\d+`)
	seeds := []int{}
	for _, res := range re.FindAllString(lines[0], -1) {
		seed, _ := strconv.Atoi(res)
		seeds = append(seeds, seed)
	}

	// fmt.Println("seeds")
	// fmt.Println(seeds)
	// results := [][]int{}

	for i, section := range lines {
		if i == 0 {
			continue
		}
		result := []int{}

		ranges := []Instruction{}
		for _, line := range strings.Split(section, "\n") {
			if strings.Contains(line, "map") {
				fmt.Println(strings.TrimSpace(line))
			} else if len(re.FindAllString(line, -1)) > 0 {
				nums := re.FindAllString(line, -1)

				dest, _ := strconv.Atoi(nums[0])
				src, _ := strconv.Atoi(nums[1])
				len, _ := strconv.Atoi(nums[2])
				ins := Instruction{dest: dest, src: src, len: len}

				ranges = append(ranges, ins)
			}
		}

		for _, seed := range seeds {
			num := seed
			// println("seed", seed)
			for _, ins := range ranges {
				if num >= ins.src && num < ins.src+ins.len {
					// fmt.Println(num, "->", num+(ins.dest-ins.src))
					num = num + (ins.dest - ins.src)
					break
				} else {
					continue
				}
			}
			result = append(result, num)
		}
		seeds = result
	}

	min := math.MaxInt
	for _, seed := range seeds {
		if seed < min {
			min = seed
		}
	}

	fmt.Println(seeds)
	fmt.Println("min", min)

}
