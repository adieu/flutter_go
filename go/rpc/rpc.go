package rpc

import (
	"bytes"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/adieu/flutter_go/go/channel/types"
)

type codec struct {
	input  *bytes.Reader
	output *bytes.Buffer
}

func (s *codec) Write(p []byte) (n int, err error) {
	return s.output.Write(p)
}

func (s *codec) Read(p []byte) (n int, err error) {
	return s.input.Read(p)
}

func (s *codec) Close() error {
	return nil
}

func newCodec(in []byte) *codec {
	return &codec{
		input:  bytes.NewReader(in),
		output: bytes.NewBuffer(nil),
	}
}

type RPCChannel struct {
	name   string
	server *rpc.Server
}

func NewRPCChannel(name string, server *rpc.Server) *RPCChannel {
	if server == nil {
		server = rpc.DefaultServer
	}
	return &RPCChannel{
		name:   name,
		server: server,
	}
}

func (r *RPCChannel) Send(message []byte, replier types.Replier) error {
	m := make([]byte, len(message))
	copy(m, message)
	go Call(r.server, m, replier)
	return nil
}

// Call a RPC function
func Call(server *rpc.Server, message []byte, replier types.Replier) {
	c := newCodec(message)
	err := server.ServeRequest(jsonrpc.NewServerCodec(c))
	if err != nil {
		replier.Reply([]byte(`{"error": "rpc error"}`))
		return
	}
	resp := c.output.Bytes()
	r := make([]byte, len(resp))
	copy(r, resp)
	replier.Reply(r)
	return
}
