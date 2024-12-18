package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func MD5String(str string, ret *string, length *int) {
	h := md5.New()
	io.WriteString(h, str)
	reta := hex.EncodeToString(h.Sum(nil))
	*ret = reta
	*length = len(hex.EncodeToString(h.Sum(nil)))
}

func MD5File(filename string, ret *string, length *int) {
	f, err := os.Open(filename)
	if err != nil {
		*ret = err.Error()
		*length = -1
		return
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		*ret = err.Error()
		*length = -2
		return
	}

	reta := hex.EncodeToString(h.Sum(nil))
	*ret = reta
	*length = len(hex.EncodeToString(h.Sum(nil)))
}

// func main() {
// 	ret := ""
// 	length := 0
// 	MD5String("../1l.json", &ret, &length)
// 	fmt.Println(ret, length)

// 	MD5File("../1l.json", &ret, &length)
// 	fmt.Println(ret, length)
// }
