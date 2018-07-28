package transfer

import (
	"io"
	"net"
	"time"
)

type TCPParam struct {
	IP      string        `param:"ip"`
	Port    string        `param:"port"`
	Timeout time.Duration `param:"timeout"`
}

type TCPTransfer struct {
	conn        net.Conn
	param       *TCPParam
	isConnected bool
}

func (tcp *TCPTransfer) Init(param interface{}) error {
	var err error
	tcp.param = param.(*TCPParam)
	_, err = tcp.connect()
	return err
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

	if err = tcp.conn.SetWriteDeadline(time.Time{}); err != nil {
		return n, err
	}
	return n, nil
}

func (tcp *TCPTransfer) Write(buf []byte) (n int, err error) {
	if err := tcp.conn.SetWriteDeadline(time.Now().Add(time.Second * 10)); err != nil {
		return 0, err
	}
	n, err = tcp.conn.Write(buf)
	if err != nil {
		return n, err
	}

	if err = tcp.conn.SetReadDeadline(time.Time{}); err != nil {
		return n, err
	}
	return n, nil
}

func (tcp *TCPTransfer) Close() error {
	return tcp.conn.Close()
}

func (tcp *TCPTransfer) connect() (bool, error) {
	var err error
	if tcp.isConnected {
		return true, err
	}

	tcp.conn, err = net.DialTimeout("tcp", tcp.param.IP+":"+ tcp.param.Port, tcp.param.Timeout)
	if err != nil {
		return false, err
	}
	tcp.isConnected = true
	return true, err
}

func (tcp *TCPTransfer)reconnect() (bool, error)  {
	tcp.conn.Close()
	tcp.isConnected = false
	return tcp.connect()
}