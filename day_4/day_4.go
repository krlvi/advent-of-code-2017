package day_4

import (
	"advent-of-code-2017/input"
	"fmt"
	"sort"
	"strings"
)

func Part_1() {
	lines := input.Lines("day_4/in.txt")
	valid := 0
	for _, l := range lines {
		if noRepeatedWords(l) {
			valid++
		}
	}
	fmt.Println(valid)
}

func Part_2() {
	lines := input.Lines("day_4/in.txt")
	valid := 0
	for _, l := range lines {
		if noAnagrams(l) {
			valid++
		}
	}
	fmt.Println(valid)
}

func noRepeatedWords(l string) bool {
	tokens := map[string]bool{}
	for _, f := range strings.Fields(l) {
		if tokens[f] == true {
			return false
		}
		tokens[f] = true
	}
	return true
}

func noAnagrams(l string) bool {
	fields := strings.Fields(l)
	for i := 0; i < len(fields); i++ {
		if i < len(fields)-1 {
			for j := i + 1; j < len(fields); j++ {
				if isAnagram(fields[i], fields[j]) {
					return false
				}
			}
		}
	}
	return true
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aSlice := strings.Split(a, "")
	bSlice := strings.Split(b, "")
	sort.Strings(aSlice)
	sort.Strings(bSlice)
	if strings.Join(aSlice, "") == strings.Join(bSlice, "") {
		return true
	}
	return false
}
