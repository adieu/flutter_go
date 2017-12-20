package registry

import (
	"errors"
	"github.com/adieu/flutter_go/go/channel/types"
	"sync"
)

var NativeManager types.Manager = nil

type Pair struct {
	Name   string
	Go     types.Channel
	Native types.Channel
}

func NewPair(name string, goChannel, nativeChannel types.Channel) *Pair {
	return &Pair{
		Name:   name,
		Go:     goChannel,
		Native: nativeChannel,
	}
}

type Registry struct {
	channelMap sync.Map
}

func NewRegistry() *Registry {
	return &Registry{}
}

var DefaultRegistry = NewRegistry()

func (r *Registry) RegisterChannel(name string, channel types.Channel) error {
	if _, ok := r.channelMap.Load(name); ok {
		return errors.New("Already registered channel")
	} else {
		if NativeManager == nil {
			r.channelMap.Store(name, NewPair(name, channel, nil))
		} else {
			n, err := NativeManager.NewChannel(name)
			if err != nil {
				return err
			}
			r.channelMap.Store(name, NewPair(name, channel, n))
		}
	}
	return nil
}

func (r *Registry) GetChannel(name string) (*Pair, error) {
	if c, ok := r.channelMap.Load(name); ok {
		return c.(*Pair), nil
	} else {
		return nil, errors.New("channel does not exist")
	}
}

func (r *Registry) GetAllChannels() map[string]*Pair {
	m := make(map[string]*Pair)
	r.channelMap.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.(*Pair)
		return true
	})
	return m
}

func Connect(name string) (types.Channel, error) {
	p, err := DefaultRegistry.GetChannel(name)
	if err != nil {
		return nil, err
	}
	if p.Native == nil {
		return nil, errors.New("Native channel is not initialized")
	}
	return p.Native, nil
}

func Listen(name string, channel types.Channel) error {
	return DefaultRegistry.RegisterChannel(name, channel)
}
