package tools

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
)

func NewTotalPage(total, pagesize int64) int {
	pages := float64(pagesize)
	totalp := float64(total)

	totalPage := math.Ceil(totalp / pages)

	return int(totalPage)
}

func MD5(name, password, salt string) string {
	saltString := "!w@k#k"
	pass := md5.New()
	io.WriteString(pass, saltString)
	io.WriteString(pass, name)
	io.WriteString(pass, salt)
	io.WriteString(pass, password)

	return fmt.Sprintf("%x", pass.Sum(nil))
}
