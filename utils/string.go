package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(value string) string {
	h := md5.New()
	io.WriteString(h, "value")
	return fmt.Sprintf("%x", h.Sum(nil))
}
