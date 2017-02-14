package manager

type Manager interface {
	Process(Message) string
	GetType() string
}