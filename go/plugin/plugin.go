package plugin

import (
	"log"

	"github.com/adieu/flutter_go/go/channel/registry"
	_ "github.com/adieu/flutter_go/go/plugin/server"
	"github.com/adieu/flutter_go/go/rpc"
)

func Bootstrap() {
	registry.Listen("flutter_go", rpc.NewRPCChannel("flutter_go", nil))
}

func init() {
	log.Print("Initializing go plugin")
}
