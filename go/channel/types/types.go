package types

type Replier interface {
	Reply([]byte) error
}

type Channel interface {
	Send([]byte, Replier) error
}

type Manager interface {
	NewChannel(string) (Channel, error)
}
