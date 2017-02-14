package manager

type Message struct {
	Command string
	Value map[string]string
}

func NewMessage(command string, value map[string]string) Message {
	return Message{command, value}
}