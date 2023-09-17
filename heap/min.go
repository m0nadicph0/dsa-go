package heap

import "math"

type MinHeap struct {
	heap []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Root() int {
	if len(h.heap) == 0 {
		return math.MaxInt
	}
	return h.heap[0]
}

func (h *MinHeap) Last() int {
	lastIndex := len(h.heap) - 1
	return h.heap[lastIndex]
}

func (h *MinHeap) Shrink() {
	lastIndex := len(h.heap) - 1
	h.heap = h.heap[:lastIndex]
}

func (h *MinHeap) Len() int {
	return len(h.heap)
}

func (h *MinHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.heapifyUp(len(h.heap) - 1)
}
func (h *MinHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}
	parentIndex := (index - 1) / 2

	if h.heap[parentIndex] > h.heap[index] {
		h.heap[parentIndex], h.heap[index] = h.heap[index], h.heap[parentIndex]
		h.heapifyUp(parentIndex)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	lci, rci := 2*index+1, 2*index+2
	minIndex := index

	if lci < h.Len() && h.heap[lci] < h.heap[minIndex] {
		minIndex = lci
	}

	if rci < h.Len() && h.heap[rci] < h.heap[minIndex] {
		minIndex = rci
	}

	if minIndex != index {
		h.heap[minIndex], h.heap[index] = h.heap[index], h.heap[minIndex]
		h.heapifyDown(minIndex)
	}
}

func (h *MinHeap) ExtractMin() int {
	if len(h.heap) == 0 {
		return math.MaxInt
	}
	min := h.Root()
	h.heap[0] = h.Last()
	h.Shrink()
	h.heapifyDown(0)
	return min
}
