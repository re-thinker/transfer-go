package transfer

import (
	"net"
	"time"
	"io"
)



type TCPTransfer struct{
	conn net.Conn
}

func (tcp *TCPTransfer) Init() error {
	var err error
	tcp.conn , err = net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}

	return nil
}

func (tcp *TCPTransfer) Read(buf []byte) (n int, err error) {
	if err := tcp.conn.SetReadDeadline(time.Now().Add(time.Second * 10)); err != nil {
		return 0, err
	}
	n, err = io.ReadFull(tcp.conn, buf)
	if err != nil {
		return n, err
	}
	/*index := 0
	try := 0
	for index < len(p) {
		n, err := tcp.conn.Read(p[index:])
		if err != nil {
			e, ok := err.(net.Error)
			if !ok || !e.Temporary() || try >= 3 {
				return index, err
			}
			try++
		}
		index += n
	}*/

	if err = tcp.conn.SetReadDeadline(time.Time{}); err != nil {
		return n, err
	}

	return n,nil
}
