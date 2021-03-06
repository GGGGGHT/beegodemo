package models

import (
	"crypto/md5"
	"fmt"
	"time"
)

func Unix2Date(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x\n", md5.Sum(data))
}
