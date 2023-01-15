package platform

// TODO
/*
Предоставляет интрфейс для управления и
Позволяет подписываться на свои события
*/
type Interface interface {
	Controls
	Subscriptions
}

type Controls interface{}

type Subscriptions interface{}
