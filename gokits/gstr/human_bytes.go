package gstr

import (
	"fmt"
	"regexp"
	"strconv"
)

// humanBytes struct
type humanBytes struct{}

const (
	_ = 1.0 << (10 * iota) // ignore first value by assigning to blank identifier
	KB
	MB
	GB
	TB
	PB
	EB
)

var (
	byteStrPattern = regexp.MustCompile(`(?i)^(-?\d+(?:\.\d+)?)\s?([KMGTPE]B?|B?)$`)
	gHumanBytes    = newHumanBytes()
)

// newhumanBytes creates a humanBytes instance.
func newHumanBytes() *humanBytes {
	return &humanBytes{}
}

// Format formats bytes integer to human readable string.
// For example, 31323 bytes will return 30.59KB.
func (*humanBytes) Format(b int64) string {
	multiple := ""
	value := float64(b)

	switch {
	case b >= EB:
		value /= EB
		multiple = "EB"
	case b >= PB:
		value /= PB
		multiple = "PB"
	case b >= TB:
		value /= TB
		multiple = "TB"
	case b >= GB:
		value /= GB
		multiple = "GB"
	case b >= MB:
		value /= MB
		multiple = "MB"
	case b >= KB:
		value /= KB
		multiple = "KB"
	case b == 0:
		return "0"
	default:
		return strconv.FormatInt(b, 10) + "B"
	}

	return fmt.Sprintf("%.2f%s", value, multiple)
}

// Parse parses human readable bytes string to bytes integer.
// For example, 6GB (6G is also valid) will return 6442450944.
func (*humanBytes) Parse(value string) (i int64, err error) {
	parts := byteStrPattern.FindStringSubmatch(value)
	if len(parts) < 3 {
		return 0, fmt.Errorf("error parsing value=%s", value)
	}
	bytesString := parts[1]
	multiple := parts[2]
	bytes, err := strconv.ParseFloat(bytesString, 64)
	if err != nil {
		return
	}

	switch multiple {
	default:
		return int64(bytes), nil
	case "K", "KB":
		return int64(bytes * KB), nil
	case "M", "MB":
		return int64(bytes * MB), nil
	case "G", "GB":
		return int64(bytes * GB), nil
	case "T", "TB":
		return int64(bytes * TB), nil
	case "P", "PB":
		return int64(bytes * PB), nil
	case "E", "EB":
		return int64(bytes * EB), nil
	}
}

// FormatHumanBytes wraps gHumanBytes Bytes's Format function.
func FormatHumanBytes(b int64) string {
	return gHumanBytes.Format(b)
}

// ParseHumanBytes wraps gHumanBytes Bytes's Parse function.
func ParseHumanBytes(val string) (int64, error) {
	return gHumanBytes.Parse(val)
}
