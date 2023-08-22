package backtrack

import "strconv"

type sVisitFn func(string)

func IPAddresses(input string) []string {
	var result []string
	forEachIPAddress(input, "", 0, 0, func(ip string) {
		result = append(result, ip)
	})
	return result
}

func forEachIPAddress(input string, currentIP string, index int, segmentNumber int, fn sVisitFn) {
	if segmentNumber == 4 && index == len(input) {
		fn(currentIP[:len(currentIP)-1])
		return
	}

	for i := 1; i <= 3; i++ {
		if index+i > len(input) {
			break
		}

		segmentStr := input[index : index+i]
		segmentInt, _ := strconv.Atoi(segmentStr)

		if segmentInt > 255 || (len(segmentStr) > 1 && segmentStr[0] == '0') {
			break
		}

		forEachIPAddress(input, currentIP+segmentStr+".", index+i, segmentNumber+1, fn)
	}
}
