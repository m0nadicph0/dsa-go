package twoptr

import "github.com/m0nadicph0/dsa-go/util"

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		area := util.Min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}
