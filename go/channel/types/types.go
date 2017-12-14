package types

type Sender interface {
	SetChannel(Channel) error
	SendMessage([]byte, Replier) error
}

type Receiver interface {
	SetChannel(Channel) error
	OnMessage([]byte, Replier) error
}

type Replier interface {
	Reply([]byte) error
}

type Channel interface {
	SendMessage([]byte, Replier) error
}

type Manager interface {
	NewSender(string) (Sender, error)
	NewReceiver(string) (Receiver, error)
}
