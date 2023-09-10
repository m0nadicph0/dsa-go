package partition_equal_subset_sum

import "github.com/m0nadicph0/dsa-go/util"

func partitionHelper(input []int, index, currentSum, totalSum int) bool {

	// Base case: If the current sum is equal to half of the total sum, we found a valid partition
	if currentSum == totalSum/2 {
		return true
	}

	// Base case: If the current index is out of bounds or the current sum exceeds half of the total sum, it's not possible
	if index == len(input) || currentSum > totalSum/2 {
		return false
	}

	return partitionHelper(input, index+1, currentSum+input[index], totalSum) || // Include
		partitionHelper(input, index+1, currentSum, totalSum) // Exclude
}

func CanPartition(input []int) bool {
	totalSum := util.Sum(input)

	// If the total sum is odd, it's impossible to partition into two equal subsets
	if totalSum%2 != 0 {
		return false
	}

	return partitionHelper(input, 0, 0, totalSum)
}
