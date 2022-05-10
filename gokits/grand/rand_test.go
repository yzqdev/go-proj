package grand

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandWithPrefix(t *testing.T) {
	rand := NewRandWithPrefix("test", 8)
	t.Log(rand)
	assert.Equal(t, 16, len(rand))
}

func TestNewRand(t *testing.T) {
	rand := NewRand(8)
	fmt.Println(rand)
	assert.Equal(t, 8, len(rand))
}

func TestRangeRand(t *testing.T) {
	for idx := 0; idx < 100; idx++ {
		rand := RangeRand(100)
		assert.True(t, 0 <= rand && rand < 100)
	}
}

func TestNormRand(t *testing.T) {
	for idx := 0; idx < 100; idx++ {
		rand := RangeRand(100)
		assert.True(t, 0 <= rand && rand < 100)
	}
}
