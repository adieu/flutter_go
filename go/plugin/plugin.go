package plugin

import (
	"log"

	_ "github.com/adieu/flutter_go/go/plugin/server"
	_ "github.com/adieu/flutter_go/go/rpc"
)

func Bootstrap() {
}

func init() {
	log.Print("Initializing go plugin")
}
