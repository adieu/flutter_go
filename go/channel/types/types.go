package types

type Replier interface {
	Reply([]byte) error
}

type Channel interface {
	SendMessage([]byte, Replier) error
}

type Manager interface {
	NewChannel(string) (Channel, error)
	SendMessage(string, []byte, Replier) error
}
