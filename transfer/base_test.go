package transfer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestNewTransfer(t *testing.T) {
	str := string("transfer=TcpTransfer;ip=127.0.0.1;port=9600;timeout=1000")
	transfer, _ := NewTransfer(str)
	tt := reflect.ValueOf(transfer)
	real := tt.Elem()
	assert.Equal(t, real.Type().String(), "transfer.TCPTransfer")
}
