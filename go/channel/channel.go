package channel

import (
	"github.com/adieu/flutter_go/go/channel/types"
	"github.com/adieu/flutter_go/go/channel/registry"
)

func RegisterSender(name string, sender types.Sender) error {
	return registry.DefaultRegistry.RegisterSender(name, sender)
}

func RegisterReceiver(name string, receiver types.Receiver) error {
	return registry.DefaultRegistry.RegisterReceiver(name, receiver)
}

func SendMessage(name string, message []byte, replier types.Replier) error {
	c, err := registry.DefaultRegistry.GetChannel(name)
	if err != nil {
		return err
	}
	err = c.SendMessage(message, replier)
	if err != nil {
		return err
	}
	return nil
}

func Init(manager types.Manager) error {
	channels := registry.DefaultRegistry.GetAllChannels()
	for name := range channels {
		c := channels[name]
		if c.Sender == nil {
			sender, err := manager.NewSender(name)
			if err != nil {
				return err
			}
			err = c.RegisterSender(sender)
			if err != nil {
				return err
			}
		}
		if c.Receiver == nil {
			receiver, err := manager.NewReceiver(name)
			if err != nil {
				return err
			}
			err = c.RegisterReceiver(receiver)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
