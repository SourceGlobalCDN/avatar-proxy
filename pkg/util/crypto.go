package util

import (
	"crypto/md5"
	"fmt"
)

func MD5(str []byte) []byte {
	m := md5.New()
	m.Write(str)
	return m.Sum(nil)
}

func MD5String(str []byte) string {
	return string(MD5(str))
}

func MD5Hex(str []byte) string {
	return fmt.Sprintf("%x", MD5(str))
}
