package heap

type KthLargest struct {
	k       int
	minHeap MinHeap
}

func NewKthLargest(k int, nums []int) KthLargest {
	minHeap := NewMinHeap()

	for _, num := range nums {
		if minHeap.Len() < k {
			minHeap.Insert(num)
		} else if num > minHeap.Root() {
			minHeap.ExtractMin()
			minHeap.Insert(num)
		}
	}

	return KthLargest{k, *minHeap}
}

func (kl *KthLargest) Add(val int) int {
	if kl.minHeap.Len() < kl.k {
		kl.minHeap.Insert(val)
	} else if val > kl.minHeap.Root() {
		kl.minHeap.ExtractMin()
		kl.minHeap.Insert(val)
	}
	return kl.minHeap.Root()
}
