package days

import (
	"log"
	"strconv"
	"strings"
)

type FrequencyChangeList struct {
	currentPosition int
	instructions []int64
}

func (fc *FrequencyChangeList) next() int64{
	position := fc.currentPosition % len(fc.instructions)
	fc.currentPosition = position + 1
	return fc.instructions[position]
}

func NewFrequencyList(instructions string) *FrequencyChangeList{
	return &FrequencyChangeList{
		instructions: parseInstructions(instructions),
	}
}

func parseInstructions(instructions string) []int64 {
	lines :=  strings.Split(instructions, "\n")
	is := make([]int64, len(lines))
	for idx, line := range lines {
		if line == "" {
			continue
		}
		numberToAdd, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		is[idx] = numberToAdd
	}
	return is
}

func Day1(instructions string) int64{
	lines :=  strings.Split(instructions, "\n")
	var total int64
	for _, line := range lines {
		if line == "" {
			return total
		}

		numberToAdd, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		total += numberToAdd
	}
	return total
}

func Day1Part2(instruction string) int64 {
	list := NewFrequencyList(instruction)
	var total, i int64
	count := make(map[int64]int64)

	for {
		test := list.next()
		if test == 0 {
			continue
		}
		total += test
		//fmt.Printf("%d: CurrentTotal: %d, Test: %d \n" , i, total, test)
		count[total]++
		if count[total] > 1 {
			//fmt.Println("already exists!", total)
			return total
		}
		i++
	}
}
