package transfer

import "io"
import (
	"snippet"
	"errors"
)

type ITransfer interface{
	io.ReadWriteCloser
	Init(param interface{}) error
}

// NewTransfer 根据传入的参数创建新的transfer
func NewTransfer(params string) (ITransfer, error) {
	var err error
	var transfer ITransfer
	mapParams := snippet.StringToMap(params)
	switch mapParams["transfer"] {
	case "ComTransfer":
		// TODO
	case "TcpTransfer":
		var tcpParam *TCPParam
		mapParams.AssignTo(&tcpParam, "param")
		transfer = &TCPTransfer{}
		err = transfer.Init(tcpParam)
	default:
		err = errors.New("unknown transfer " + params)
	}

	return transfer, err
}