package tools

import "math"

func NewTotalPage(total, pagesize int64) int {
	pages := float64(pagesize)
	totalp := float64(total)

	totalPage := math.Ceil(totalp / pages)

	return int(totalPage)
}
