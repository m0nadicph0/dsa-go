package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthLargest(t *testing.T) {
	k := 3
	stream := []int{4, 5, 8, 2}
	kthLargest := NewKthLargest(k, stream)

	assert.Equal(t, 4, kthLargest.Add(3))
	assert.Equal(t, 5, kthLargest.Add(5))
	assert.Equal(t, 5, kthLargest.Add(10))
	assert.Equal(t, 8, kthLargest.Add(9))
	assert.Equal(t, 8, kthLargest.Add(4))

}
