package days

import (
	"github.com/steakknife/hamming"
	"strconv"
	"strings"
)

func Day2(input string) string {
	lines := strings.Split(input, "\n")
	threes := 0
	twos := 0

	for _, line := range lines {
		//fmt.Println(line)
		count := letterCount(line)
		var two, three bool
		for _, v := range count {
			two = two || v == 2
			three = three || v == 3
			//fmt.Println(line, two, three)
		}
		if two {
			twos += 1
		}
		if three {
			threes += 1
		}
	}
	return strconv.Itoa(threes * twos)
}

func Day2Part2(input string) string {
	lines := strings.Split(input, "\n")
	lowestHamming := 100000000
	bestCandidates := []string{"", ""}

	for i:=0; i<len(lines); i++{
		for j:=0; j<len(lines); j++ {
			if i==j {
				continue
			}
			ham := hamming.Strings(lines[i], lines[j])
			if ham <= lowestHamming {
				lowestHamming = ham
				bestCandidates[0] = lines[i]
				bestCandidates[1] = lines[j]
			}
		}
	}
	return commonString(bestCandidates[0], bestCandidates[1])
}

// I wanted to do set union.  But I also wanted to get quick solution :)
func commonString(b0, b1 string) string {
	if len(b0) != len(b1) {
		return ""
	}
	var result []byte
	for i, c := range b0 {
		if b1[i] == uint8(c) {
			result = append(result, uint8(c))
		}
	}
	return string(result)
}

func letterCount(input string) map[rune]int{
	count := map[rune]int{}
	for _, c := range input {
		count[c]++
	}
	return count
}
