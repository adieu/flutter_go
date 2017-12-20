package channel

import (
	"github.com/adieu/flutter_go/go/channel/registry"
	"github.com/adieu/flutter_go/go/channel/types"
)

func Send(name string, message []byte, replier types.Replier) error {
	p, err := registry.DefaultRegistry.GetChannel(name)
	if err != nil {
		return err
	}
	err = p.Go.Send(message, replier)
	if err != nil {
		return err
	}
	return nil
}

func Connect(name string) (types.Channel, error) {
	p, err := registry.DefaultRegistry.GetChannel(name)
	if err != nil {
		return nil, err
	}
	return p.Go, nil
}

func Init(manager types.Manager) error {
	registry.NativeManager = manager
	pairs := registry.DefaultRegistry.GetAllChannels()
	for name := range pairs {
		c := pairs[name]
		if c.Native == nil {
			channel, err := manager.NewChannel(name)
			if err != nil {
				return err
			}
			c.Native = channel
		}
	}
	return nil
}
