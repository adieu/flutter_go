package rpc

import (
	"bytes"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/adieu/flutter_go/go/channel/registry"
	"github.com/adieu/flutter_go/go/channel/types"
)

type stream struct {
	input  *bytes.Reader
	output *bytes.Buffer
}

func (s *stream) Write(p []byte) (n int, err error) {
	return s.output.Write(p)
}

func (s *stream) Read(p []byte) (n int, err error) {
	return s.input.Read(p)
}

func (s *stream) Close() error {
	return nil
}

func newStream(in []byte) *stream {
	return &stream{
		input:  bytes.NewReader(in),
		output: bytes.NewBuffer(nil),
	}
}

type RPCReceiver struct {
}

func (r *RPCReceiver) OnMessage(message []byte, replier types.Replier) error {
	m := make([]byte, len(message))
	copy(m, message)
	go Call(m, replier)
	return nil
}

func (r *RPCReceiver) SetChannel(channel types.Channel) error {
	return nil
}

// Call a RPC function
func Call(message []byte, replier types.Replier) {
	s := newStream(message)
	err := rpc.ServeRequest(jsonrpc.NewServerCodec(s))
	if err != nil {
		replier.Reply([]byte(`{"error": "rpc error"}`))
		return
	}
	resp := s.output.Bytes()
	r := make([]byte, len(resp))
	copy(r, resp)
	replier.Reply(r)
	return
}

func init() {
	registry.DefaultRegistry.RegisterReceiver("go_rpc", &RPCReceiver{})
}
