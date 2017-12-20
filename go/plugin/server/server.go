package server

import (
	"net/rpc"
)

type Builtin struct{}

func (*Builtin) GetPlatformVersion(_ *struct{}, reply *string) error {
	*reply = "Flutter Go"
	return nil
}

func init() {
	rpc.Register(new(Builtin))
}
