package main

import "fmt"

func main() {
	RecPermute("", "")
}

func RecPermute(soFar string, rest string) {
	fmt.Println("soFar:", soFar, ", rest:", rest)
	if len(rest) == 0 {
		fmt.Println(soFar)
		return
	}

	for i := 0; i < len(rest); i++ {
		next := soFar + string(rest[i])
		remaining := string(rest[0:i]) + string(rest[i+1:])
		RecPermute(next, remaining)
	}
}
