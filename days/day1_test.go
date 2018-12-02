package days

import (
	"github.com/tlcowling/aoc2018/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay1(t *testing.T) {
	cases := map[string]int64 {
		"../inputs/day1input_test": 1,
		"../inputs/day1input_test2": 4,
	}

	for input, expected := range cases {
		s := utils.ReadFileAsString(input)
		actual := Day1(s)
		assert.Equal(t, expected, actual, "Frequency list total")
	}
}

func TestDay1_2(t *testing.T) {
	cases := map[string]int64 {
		"../inputs/day1input_test": 14,
		"../inputs/day1input_test2": 5,
	}

	for input, expected := range cases {
		s := utils.ReadFileAsString(input)
		actual := Day1Part2(s)
		assert.Equal(t, expected, actual, "Frequency list total")
	}
}