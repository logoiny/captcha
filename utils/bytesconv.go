package utils

import (
	"reflect"
	"unsafe"
)

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}


func DigitsToString(digits []byte)string{
	b := make([]byte, len(digits))

	for i, _ := range digits {
		b[i] = digits[i] + 48
	}
	return  BytesToString(b)
}

func DigitsToByte(digits string)[]byte{
	b := make([]byte, len(digits))

	for i, _ := range digits {
		b[i] = digits[i] - 48
	}
	return  b
}