package demo1

func Call_Send5() {
	m := NewMessageService()
	m.Pass("hello", "Tom")
}

func Call_Send6() {

	m := NewMessageService()
	m.Pass("hello", "Tom")
}

func Call_Send7() {
	m := NewMessageService()
	m.Pass("hello", "Tom")
}

func Call_Send8() {
	m := NewMessageService()
	m.Pass("hello", "Tom")
}

type IMessageService interface {
	Pass(msg string, customer string)
}

type Test1 interface {
	Cal()
}
