package twoptr

import "github.com/m0nadicph0/dsa-go/util"

func SortColors(colors []int) []int {
	red, white, blue := 0, 0, len(colors)-1
	for white <= blue {
		if colors[white] == 0 {
			if colors[red] != 0 {
				util.Swap(colors, white, red)
			}
			red++
			white++
		} else if colors[white] == 1 {
			white++
		} else {
			if colors[blue] != 2 {
				util.Swap(colors, white, blue)
			}
			blue--
		}
	}
	return colors
}
