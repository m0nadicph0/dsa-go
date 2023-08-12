package recursion

import (
	"fmt"
	"io"
)

func Print(w io.Writer, a []int) {
	if len(a) == 0 {
		return
	}
	fmt.Fprint(w, a[0])
	Print(w, a[1:])
}

func ReversePrint(w io.Writer, a []int) {
	if len(a) == 0 {
		return
	}
	ReversePrint(w, a[1:])
	fmt.Fprint(w, a[0])
}

func ReversePrintAlt(w io.Writer, a []int) {
	if len(a) == 0 {
		return
	}
	fmt.Fprint(w, a[len(a)-1])
	ReversePrint(w, a[0:len(a)-1])
}

func Length(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return Length(a[1:]) + 1
}

func Replace(input []int, target, replacement int) {
	if len(input) == 0 {
		return
	}

	if input[0] == target {
		input[0] = replacement
	}
	Replace(input[1:], target, replacement)
}
