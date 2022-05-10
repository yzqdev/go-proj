package gstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.True(t, IsDigit("12345"))
	assert.False(t, IsDigit("a12345"))
	assert.False(t, IsDigit("12345b"))
}

func TestIsChinese(t *testing.T) {
	assert.True(t, IsChinese("测试aaa"))
	assert.False(t, IsDigit("测试a"))
}

func TestIsLetterNumUnline(t *testing.T) {
	assert.True(t, IsLetterNumUnline("Abcd1234_"))
	assert.False(t, IsLetterNumUnline("测试a"))
}

func TestIsChineseLetterNumUnline(t *testing.T) {
	assert.True(t, IsChineseLetterNumUnline("测试Abcd1234_"))
}

func TestIsEmail(t *testing.T) {
	assert.True(t, IsEmail("test@gmail.com"))
	assert.True(t, IsEmail("test2@gmail.com"))
	assert.True(t, IsEmail("test_2@gmail.com"))
	assert.True(t, IsEmail("1-test_2@gmail-ba1.com"))
}

func TestIsURL(t *testing.T) {
	assert.True(t, IsURL("https://github.com/xtfly/gokit"))
	assert.True(t, IsURL("https://github.com/xtfly/gokit?a=b"))
	assert.True(t, IsURL("https://github.com/xtfly/gokit#命名约定"))
	assert.True(t, IsURL("https://github.com/xtfly/gokit?a=1&wd=%20新闻&oq=%25E6%2596%25B0%25E9%2597%25BB"))
}
