package utils

import (
	"regexp"
)

func FindMatches(pattern string, search string) ([]string, []int) {
	regex := regexp.MustCompile(pattern)
	segments := regex.FindAllStringSubmatchIndex(search, -1)

	breakpoints := []int{0}
	for _, segment := range segments {
		breakpoints = append(breakpoints, segment[0], segment[1])
	}
	breakpoints = append(breakpoints, len(search))

	substrings := []string{}
	matchIndices := []int{}

	for i := 0; i < len(breakpoints)-1; i++ {
		start := breakpoints[i]
		end := breakpoints[i+1]
		substring := search[start:end]
		substrings = append(substrings, substring)
		if regex.MatchString(substring) {
			matchIndices = append(matchIndices, len(substrings)-1)
		}
	}

	return substrings, matchIndices
}
