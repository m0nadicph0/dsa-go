package lis

import (
	"github.com/m0nadicph0/dsa-go/util"
	"math"
)

func LengthOfLongestIncSubsequence(input []int) int {
	return lisHelper(input, math.MinInt64, 0)
}

func lisHelper(input []int, previous int, index int) int {
	if index == len(input) {
		return 0
	}

	include := 0
	if input[index] > previous {
		include = 1 + lisHelper(input, input[index], index+1)
	}

	exclude := lisHelper(input, previous, index+1)

	return util.Max(include, exclude)
}
