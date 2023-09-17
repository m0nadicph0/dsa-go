package heap

func FindKthSmallest(nums []int, k int) int {
	mh := NewMinHeap()
	for _, num := range nums {
		mh.Insert(num)
	}
	for i := 1; i < k; i++ {
		mh.ExtractMin()
	}
	return mh.Root()
}
