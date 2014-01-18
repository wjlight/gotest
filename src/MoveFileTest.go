package main

import (
	"fmt"

	"syscall"
	"unicode/utf16"

	"io"
	"os"
)

//字符串转换来unit16
func StringToUTF16(s string) []uint16 {
	return utf16.Encode([]rune(s + "\x00"))
}

func Move(frompath string, topath string) {
	from := &StringToUTF16(frompath)[0]
	to := &StringToUTF16(topath)[0]
	err := syscall.MoveFile(from, to)

	fmt.Println(err)
}

func CopyFile(src, dst string) (int64, error) {
	sf, err := os.Open(src)
	if err != nil {
		fmt.Printf("err", err)
		return 0, err
	}
	file, _ := os.Create(dst)
	return io.Copy(file, sf)
}

func main() {
	src := "E:/goWork/helloworld.exe"
	dst := "E:/helloworld.exe"
	// move(src, dst)
	// src := "E:/goWork/helloworld.go"
	// dst := "E:/g"
	CopyFile(src, dst)
}
