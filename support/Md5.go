package support

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func GeneratorMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}
