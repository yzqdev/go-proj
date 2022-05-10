package page

import (
	"math"
)

func CalculateTotalPage(totalCount, pageSize float64) int {
	totalPage := totalCount / pageSize

	return int(math.Ceil(totalPage))
}
