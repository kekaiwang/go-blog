package tools

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"strconv"
)

// NewTotalPage. new page
func NewTotalPage(total, pagesize int64) int {
	pages := float64(pagesize)
	totalp := float64(total)

	totalPage := math.Ceil(totalp / pages)

	return int(totalPage)
}

// MD5. md5 func
func MD5(name, password, salt string) string {
	// salt
	saltString := "!w@k#k"
	pass := md5.New()
	io.WriteString(pass, saltString)
	io.WriteString(pass, name)
	io.WriteString(pass, salt)
	io.WriteString(pass, password)

	return fmt.Sprintf("%x", pass.Sum(nil))
}

// NewLimitOffset. build page size
func NewLimitOffset(limitStr, pageStr string) (int64, int64, int64) {
	var (
		offset int64
		page   int64
		limit  int64
	)

	page, _ = strconv.ParseInt(pageStr, 10, 64)
	limit, _ = strconv.ParseInt(limitStr, 10, 64)

	if page == 0 {
		page = 1
	}

	offset = (page - 1) * limit

	return offset, page, limit
}
