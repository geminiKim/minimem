package store

type Manager interface {
	Process(Message) string
	GetType() string
}