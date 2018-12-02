package main

import (
	"fmt"
	"github.com/tlcowling/aoc2018/days"
	"github.com/tlcowling/aoc2018/utils"
)

func main() {
	fmt.Println(days.Day1(utils.ReadFileAsString("./inputs/day1input")))
	fmt.Println(days.Day1Part2(utils.ReadFileAsString("./inputs/day1input")))
	fmt.Println(days.Day2(utils.ReadFileAsString("./inputs/day2input")))
	fmt.Println(days.Day2Part2(utils.ReadFileAsString("./inputs/day2input")))
}
