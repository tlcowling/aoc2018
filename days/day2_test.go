package days

import (
	"github.com/stretchr/testify/assert"
	"github.com/tlcowling/aoc2018/utils"
	"testing"
)

func TestDay2_1(t *testing.T) {
	cases := map[string]string {
		"../inputs/day2input_test": "12",
	}

	for input, expected := range cases {
		s := utils.ReadFileAsString(input)
		actual := Day2(s)
		assert.Equal(t, expected, actual)
	}
}

func TestDay2_2(t *testing.T) {
	cases := map[string]string {
		"../inputs/day2input_test2": "fgij",
	}

	for input, expected := range cases {
		s := utils.ReadFileAsString(input)
		actual := Day2Part2(s)
		assert.Equal(t, expected, actual)
	}
}