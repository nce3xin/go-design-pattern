package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList_Iterator(t *testing.T) {
	list := ArrayList{"a", "b", "c"}
	iterator := list.Iterator()
	i := 0
	for iterator.HasNext() {
		assert.Equal(t, list[i], iterator.CurrentItem())
		iterator.Next()
		i++
	}
}
