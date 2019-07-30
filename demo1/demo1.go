package demo1

type MessageService struct {
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (m *MessageService) Pass(msg string, customer string) {
	// some code
}

func Pass() {}

//Call_Send1
func Call_Send1() {
	m := NewMessageService()
	m.Pass("hello", "Tom")
}

//Call_Send2
func Call_Send2() {
	m := NewMessageService()
	m.Pass("hello", "")
}

func Call_Send3() {
	m := NewMessageService()
	m.Pass("hello", "")
}

func Call_Send4() {
	m := NewMessageService()
	m.Pass("hello", "")
}
