package heap

import "math"

type MaxHeap struct {
	heap []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.heapifyUp(len(h.heap) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}
	parentIndex := (index - 1) / 2
	if h.heap[index] > h.heap[parentIndex] {
		h.heap[index], h.heap[parentIndex] = h.heap[parentIndex], h.heap[index]
		h.heapifyUp(parentIndex)
	}
}

func (h *MaxHeap) ExtractMax() int {
	if len(h.heap) == 0 {
		return -1 * math.MaxInt
	}
	max := h.Root()
	h.heap[0] = h.Last()
	h.Shrink()
	h.heapifyDown(0)
	return max
}

func (h *MaxHeap) Root() int {
	return h.heap[0]
}

func (h *MaxHeap) Last() int {
	return h.heap[(len(h.heap) - 1)]
}

func (h *MaxHeap) Shrink() {
	lastIndex := len(h.heap) - 1
	h.heap = h.heap[:lastIndex]
}

func (h *MaxHeap) heapifyDown(index int) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2
	maxIndex := index

	if leftChildIndex < len(h.heap) && h.heap[leftChildIndex] > h.heap[maxIndex] {
		maxIndex = leftChildIndex
	}

	if rightChildIndex < len(h.heap) && h.heap[rightChildIndex] > h.heap[maxIndex] {
		maxIndex = rightChildIndex
	}

	if maxIndex != index {
		h.heap[maxIndex], h.heap[index] = h.heap[index], h.heap[maxIndex]
		h.heapifyDown(maxIndex)
	}
}
