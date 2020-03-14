package common

import (
	"crypto/md5"
	"fmt"
)

func Md5(data string) string {
	return fmt.Sprintf("%x\n", md5.Sum([]byte(data)))
}
