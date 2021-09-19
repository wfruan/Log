package main

import (
	"crypto/md5"
	"fmt"
)

func main()  {
	data := []byte("Hello,World!")
	hashValue := md5.Sum(data)
	fmt.Println(hashValue)
	fmt.Printf("原始值: %s\n", data)
	fmt.Printf("MD5值(binary): %b\n", hashValue)
	fmt.Printf("MD5值: %c\n", hashValue)
	fmt.Printf("MD5值(xo): %x\n", hashValue)
	fmt.Printf("MD5值: %d\n", hashValue)
}
