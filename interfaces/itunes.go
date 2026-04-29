package interfaces

type ITunes interface {
	Initiate() error
	Close() error
}
