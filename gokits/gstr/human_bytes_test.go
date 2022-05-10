package gstr

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesFormatHumanBytes(t *testing.T) {
	// B
	b := FormatHumanBytes(0)
	assert.Equal(t, "0", b)
	// B
	b = FormatHumanBytes(515)
	assert.Equal(t, "515B", b)

	// KB
	b = FormatHumanBytes(31323)
	assert.Equal(t, "30.59KB", b)

	// MB
	b = FormatHumanBytes(13231323)
	assert.Equal(t, "12.62MB", b)

	// GB
	b = FormatHumanBytes(7323232398)
	assert.Equal(t, "6.82GB", b)

	// TB
	b = FormatHumanBytes(7323232398434)
	assert.Equal(t, "6.66TB", b)

	// PB
	b = FormatHumanBytes(9923232398434432)
	assert.Equal(t, "8.81PB", b)

	// EB
	b = FormatHumanBytes(math.MaxInt64)
	assert.Equal(t, "8.00EB", b)
}

func TestBytesParseHumanBytesErrors(t *testing.T) {
	_, err := ParseHumanBytes("B999")
	if assert.Error(t, err) {
		assert.EqualError(t, err, "error parsing value=B999")
	}
}

func TestFloats(t *testing.T) {
	// From string:
	str := "12.25KB"
	value, err := ParseHumanBytes(str)
	assert.NoError(t, err)
	assert.Equal(t, int64(12544), value)

	str2 := FormatHumanBytes(value)
	assert.Equal(t, str, str2)

	// To string:
	val := int64(13233029)
	str = FormatHumanBytes(val)
	assert.Equal(t, "12.62MB", str)

	val2, err := ParseHumanBytes(str)
	assert.NoError(t, err)
	assert.Equal(t, val, val2)
}

func TestBytesParseHumanBytes(t *testing.T) {
	// B
	b, err := ParseHumanBytes("999")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(999), b)
	}
	b, err = ParseHumanBytes("-100")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(-100), b)
	}
	b, err = ParseHumanBytes("100.1")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(100), b)
	}
	b, err = ParseHumanBytes("515B")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(515), b)
	}

	// B with space
	b, err = ParseHumanBytes("515 B")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(515), b)
	}

	// KB
	b, err = ParseHumanBytes("12.25KB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12544), b)
	}
	b, err = ParseHumanBytes("12KB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12288), b)
	}
	b, err = ParseHumanBytes("12K")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12288), b)
	}

	// KB with space
	b, err = ParseHumanBytes("12.25 KB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12544), b)
	}
	b, err = ParseHumanBytes("12 KB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12288), b)
	}
	b, err = ParseHumanBytes("12 K")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12288), b)
	}

	// MB
	b, err = ParseHumanBytes("2MB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(2097152), b)
	}
	b, err = ParseHumanBytes("2M")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(2097152), b)
	}

	// GB with space
	b, err = ParseHumanBytes("6 GB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(6442450944), b)
	}
	b, err = ParseHumanBytes("6 G")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(6442450944), b)
	}

	// GB
	b, err = ParseHumanBytes("6GB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(6442450944), b)
	}
	b, err = ParseHumanBytes("6G")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(6442450944), b)
	}

	// TB
	b, err = ParseHumanBytes("5TB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(5497558138880), b)
	}
	b, err = ParseHumanBytes("5T")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(5497558138880), b)
	}

	// TB with space
	b, err = ParseHumanBytes("5 TB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(5497558138880), b)
	}
	b, err = ParseHumanBytes("5 T")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(5497558138880), b)
	}

	// PB
	b, err = ParseHumanBytes("9PB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10133099161583616), b)
	}
	b, err = ParseHumanBytes("9P")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10133099161583616), b)
	}

	// PB with space
	b, err = ParseHumanBytes("9 PB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10133099161583616), b)
	}
	b, err = ParseHumanBytes("9 P")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10133099161583616), b)
	}

	// EB
	b, err = ParseHumanBytes("8EB")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b-1)
	}
	b, err = ParseHumanBytes("8E")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b-1)
	}

	// EB with spaces
	b, err = ParseHumanBytes("8 EB")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b-1)
	}
	b, err = ParseHumanBytes("8 E")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b-1)
	}
}