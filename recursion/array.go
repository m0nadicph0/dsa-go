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

func RemoveAll(arr []int, target int) []int {
	if len(arr) == 0 {
		return arr
	}

	if arr[0] == target {
		return RemoveAll(arr[1:], target)
	}
	return append([]int{arr[0]}, RemoveAll(arr[1:], target)...)
}

func RemoveConsecutive(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if arr[0] == arr[1] {
		return RemoveConsecutive(arr[1:])
	}
	return append([]int{arr[0]}, RemoveConsecutive(arr[1:])...)
}

func GetAllSubsequences(input string) []string {
	result := make([]string, 0)
	forEachSubsequence(input, "", func(s string) {
		result = append(result, s)
	})
	return result
}

func forEachSubsequence(s string, output string, fm func(s string)) {
	if len(s) == 0 {
		fm(output)
		return
	}
	forEachSubsequence(s[1:], output, fm)
	forEachSubsequence(s[1:], output+string(s[0]), fm)
}

func StrToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	last := int(s[len(s)-1]) - 48
	return StrToInt(s[0:len(s)-1])*10 + last
}

func AllPermutation(input string) []string {
	result := make([]string, 0)
	forEachPermutation([]byte(input), 0, len(input)-1, func(s string) {
		result = append(result, s)
	})
	return result
}

func forEachPermutation(input []byte, start int, end int, fn func(s string)) {
	if start == end {
		fn(string(input))
		return
	}

	for i := start; i <= end; i++ {
		swap(input, start, i)
		forEachPermutation(input, start+1, end, fn)
		swap(input, start, i)
	}
}

func swap(input []byte, i int, j int) {
	tmp := input[i]
	input[i] = input[j]
	input[j] = tmp
}

func AllPermutationAlt(input string) []string {
	result := make([]string, 0)
	forEachPermutationAlt("", input, func(permutation string) {
		result = append(result, permutation)
	})
	return result
}

func forEachPermutationAlt(soFar string, rest string, fn func(string)) {
	if len(rest) == 0 {
		if len(soFar) > 0 {
			fn(soFar)
		}
		return
	}

	for i := 0; i < len(rest); i++ {
		next := soFar + string(rest[i])
		remaining := string(rest[0:i]) + string(rest[i+1:])
		forEachPermutationAlt(next, remaining, fn)
	}
}
