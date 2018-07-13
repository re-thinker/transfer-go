package convert

import (
	"bytes"
	"encoding/binary"
)

func Bytetoint(by []byte, num *int32)  {
	buf := bytes.NewBuffer(by)
	binary.Read(buf, binary.BigEndian, num)
}

func Inttobyte(num *int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}

