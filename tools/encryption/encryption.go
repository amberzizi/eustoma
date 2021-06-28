// @Title   encryption.go
// @Description  加解密
// @Author  amberhu  20210624
// @Update
package encryption

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(n string) string {
	w := md5.New()
	io.WriteString(w, n)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
