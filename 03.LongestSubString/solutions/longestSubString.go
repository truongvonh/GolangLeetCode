package solutions

import (
	"fmt"
	"math"
)

const DefaultLength = 1

func LongestSubString(sentences string) float64 {
	var longest float64 = 0
	var start float64 = 0
	mapObj := make(map[string]float64)

	for i := 0; i < len(sentences); i++ {
		characterAtIndex := string(sentences[i])
		if valueOfChar, ok := mapObj[characterAtIndex]; ok && start <= valueOfChar {
			longest = math.Max(float64(i)-start, longest)
			fmt.Printf("Longest by index %v has max value %v\n", i, longest)
			start = mapObj[characterAtIndex] + 1
		}

		mapObj[characterAtIndex] = float64(i)
	}
	fmt.Printf("Map object store data is %v\n", mapObj)

	return math.Max(float64(len(sentences))-start, longest)
}
