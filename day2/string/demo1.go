package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

//[]byte(s)转换模拟实现
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

//string(bytes)转换模拟实现
func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)

	return p
}

//[]rune(s)转换模拟实现
func str2runes(s []byte) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

//string(runes)转换模拟实现
func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}

func main() {
	fmt.Printf("%#v\n", []rune("世界"))              // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界
}
