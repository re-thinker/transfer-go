package transfer

import "io"

type ITransfer interface{
	io.ReadWriteCloser
	Init(param interface{}) error
}

