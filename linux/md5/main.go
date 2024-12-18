package main

import (
	"C"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func MD5String(str string, ret *string, length *int) {
	h := md5.New()
	io.WriteString(h, str)
	var reta string
	reta = hex.EncodeToString(h.Sum(nil))
	*ret = reta
	*length = len(hex.EncodeToString(h.Sum(nil)))
	return
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

	var reta string
	reta = hex.EncodeToString(h.Sum(nil))
	*ret = reta
	*length = len(hex.EncodeToString(h.Sum(nil)))
	return
}

func main() {
	ret := ""
	length := 0
	MD5String("m.png", &ret, &length)
	fmt.Println(ret, length)

	MD5File("m.png", &ret, &length)
	fmt.Println(ret, length)
}
