package sliding_window

func MaxSlidingWindow(nums []int, w int) []int {
	var result []int
	window := NewDeque()

	for i := 0; i < len(nums); i++ {

		// Remove indices outside the current window
		for window.Len() > 0 && window.Front() <= i-w {
			window.PopFront()
		}

		// Remove indices of smaller elements as they won't contribute to the maximum
		for window.Len() > 0 && nums[i] >= nums[window.Back()] {
			window.PopBack()
		}

		// Add the current index to the deque
		window.PushBack(i)

		// Add the maximum value to the result for each valid window
		if i >= w-1 {
			result = append(result, nums[window.Front()])
		}
	}

	return result
}
