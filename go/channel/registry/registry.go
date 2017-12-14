package registry

import (
	"errors"
	"sync"
	"github.com/adieu/flutter_go/go/channel/types"
)

type Channel struct {
	Name string
	Sender types.Sender
	Receiver types.Receiver
}

func NewChannel(name string) *Channel {
	return &Channel{
		Name: name,
	}
}

func NewSenderChannel(name string, sender types.Sender) *Channel {
	c := &Channel{
		Name: name,
		Sender: sender,
	}
	sender.SetChannel(c)
	return c
}

func NewReceiverChannel(name string, receiver types.Receiver) *Channel {
	c := &Channel{
		Name: name,
		Receiver: receiver,
	}
	receiver.SetChannel(c)
	return c
}

func (c *Channel) RegisterSender(sender types.Sender) error {
	c.Sender = sender
	sender.SetChannel(c)
	return nil
}

func (c *Channel) RegisterReceiver(receiver types.Receiver) error {
	c.Receiver = receiver
	receiver.SetChannel(c)
	return nil
}

func (c *Channel) SendMessage(message []byte, replier types.Replier) error {
	if c.Receiver == nil {
		return nil
	}
	return c.Receiver.OnMessage(message, replier)
}

type Registry struct {
	channelMap sync.Map
}

func NewRegistry() *Registry {
	return &Registry{}
}

var DefaultRegistry = NewRegistry()

func (r *Registry) RegisterSender(name string, sender types.Sender) error {
	if c, ok := r.channelMap.Load(name); ok {
		if c.(*Channel).Sender == nil {
			return c.(*Channel).RegisterSender(sender)
		} else {
			return errors.New("")
		}
	} else {
		r.channelMap.Store(name, NewSenderChannel(name, sender))
	}
	return nil
}

func (r *Registry) RegisterReceiver(name string, receiver types.Receiver) error {
	if c, ok := r.channelMap.Load(name); ok {
		if c.(*Channel).Receiver == nil {
			return c.(*Channel).RegisterReceiver(receiver)
		} else {
			return errors.New("")
		}
	} else {
		r.channelMap.Store(name, NewReceiverChannel(name, receiver))
	}
	return nil
}

func (r *Registry) GetChannel(name string) (*Channel, error) {
	if c, ok := r.channelMap.Load(name); ok {
		return c.(*Channel), nil
	} else {
		return nil, errors.New("channel does not exist")
	}
}

func (r *Registry) GetAllChannels() map[string]*Channel {
	m := make(map[string]*Channel)
	r.channelMap.Range(func (key, value interface{}) bool{
		m[key.(string)] = value.(*Channel)
		return true
        })
	return m
}
