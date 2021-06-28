// @Title   randstring.go
// @Description  随机字符串
// @Author  amberhu  20210624
// @Update
package randstring

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	randfinal := rand.New(rand.NewSource(time.Now().UnixNano()))
	bstring := make([]rune, n)
	for i := range bstring {
		bstring[i] = letters[randfinal.Intn(len(letters))]
	}
	return string(bstring)
}
