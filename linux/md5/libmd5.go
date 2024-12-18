package main

import (
	"C"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

/*MD5String is a func to calculate MD5 of a string*/
//export MD5String
func MD5String(str *C.char, ret **C.char, length *int) {
	strl := C.GoString(str)
	strl = strings.TrimSpace(strl)
	h := md5.New()
	io.WriteString(h, strl)
	var reta string
	reta = hex.EncodeToString(h.Sum(nil))
	*ret = C.CString(reta)
	*length = len(hex.EncodeToString(h.Sum(nil)))
	return
}

/*MD5File is a func to calculate MD5 of a file*/
//export MD5File
func MD5File(filename *C.char, ret **C.char, length *int) {
	filenamel := C.GoString(filename)
	filenamel = strings.TrimSpace(filenamel)
	f, err := os.Open(filenamel)
	if err != nil {
		*ret = C.CString(err.Error())
		*length = -1
		return
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		*ret = C.CString(err.Error())
		*length = -2
		return
	}

	var reta string
	reta = hex.EncodeToString(h.Sum(nil))
	*ret = C.CString(reta)
	*length = len(hex.EncodeToString(h.Sum(nil)))
	return
}

func main() {}
