package main

import (
	"convert"
	"fmt"
	"transfer"
)

func main()  {
	by := []byte{0x00, 0x00, 0x03, 0xe8}
	var num int32
	convert.Bytetoint(by, &num)
	fmt.Println("num :", num)

	by2 := []byte{}
	var num2 int32
	num2 = 333
	by2 = convert.Inttobyte(&num2)
	fmt.Println(by2)

	var num3 int32
	convert.Bytetoint(by2, &num3)
	fmt.Println(num3)

	tcp := transfer.TCPTransfer{}
	tcp.Init()
	buf := make([]byte, 10)
	tcp.Read(buf)
	fmt.Println(buf)

}