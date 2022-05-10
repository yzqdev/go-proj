package gstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinInts(t *testing.T) {
	// test empty slice
	is := []int64{}
	s := JoinInts(is)
	assert.Equal(t, "", s)

	// test len(slice)==1
	is = []int64{1}
	s = JoinInts(is)
	assert.Equal(t, "1", s)

	// test len(slice)>1
	is = []int64{1, 2, 3}
	s = JoinInts(is)
	assert.Equal(t, "1,2,3", s)
}

func TestSplitInts(t *testing.T) {
	// test empty slice
	s := ""
	is, err := SplitInts(s)
	assert.Equal(t, 0, len(is))
	assert.NoError(t, err)

	// test split int64
	s = "1,2,3"
	is, err = SplitInts(s)
	assert.Equal(t, 3, len(is))
	assert.NoError(t, err)
}

func BenchmarkJoinInts(b *testing.B) {
	is := make([]int64, 10000, 10000)
	for i := int64(0); i < 10000; i++ {
		is[i] = i
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			JoinInts(is)
		}
	})
}
