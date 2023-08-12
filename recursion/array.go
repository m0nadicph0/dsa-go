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

func IsSorted(a []int) bool {
	if len(a) <= 1 {
		return true
	}
	return (a[0] < a[1]) && IsSorted(a[1:])
}

func Sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return a[0] + Sum(a[1:])
}

func IsPresent(a []int, target int) bool {
	if len(a) == 0 {
		return false
	}
	return (a[0] == target) || IsPresent(a[1:], target)
}

func FirstIndexOf(a []int, target int) int {
	return firstIndex(a, target, 0)
}
func firstIndex(a []int, target int, start int) int {
	if len(a) == 0 {
		return -1
	}
	if a[0] == target {
		return start
	}
	return firstIndex(a[1:], target, start+1)
}

func LastIndexOf(a []int, target int) int {
	return lastIndex(a, target, len(a)-1)
}
func lastIndex(a []int, target int, end int) int {
	last := len(a) - 1
	if len(a) == 0 {
		return -1
	}
	if a[last] == target {
		return end
	}
	return lastIndex(a[0:last], target, end-1)
}

func PrintAllOccurrence(w io.Writer, a []int, target int) {
	printAllOccurrenceRec(w, a, target, 0)
}

func printAllOccurrenceRec(w io.Writer, a []int, target int, index int) {
	if len(a) == 0 {
		return
	}
	if a[0] == target {
		fmt.Fprint(w, index)
	}
	printAllOccurrenceRec(w, a[1:], target, index+1)
}

func CountOccurrence(a []int, target int) int {
	if len(a) == 0 {
		return 0
	}

	if a[0] == target {
		return CountOccurrence(a[1:], target) + 1
	} else {
		return CountOccurrence(a[1:], target) + 0
	}
}

func StoreAllOccurrence(a []int, target int) []int {
	acct := make([]int, 0)
	storeAllOccurrenceRec(a, target, 0, func(index int) {
		acct = append(acct, index)
	})
	return acct
}

func storeAllOccurrenceRec(a []int, target int, index int, fn func(int)) {
	if len(a) == 0 {
		return
	}
	if a[0] == target {
		fn(index)
	}

	storeAllOccurrenceRec(a[1:], target, index+1, fn)
}

func IsPalindrome(a []int) bool {
	if len(a) <= 1 {
		return true
	}
	return (a[0] == a[len(a)-1]) && IsPalindrome(a[1:len(a)-1])
}
