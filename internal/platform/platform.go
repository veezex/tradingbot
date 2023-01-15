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

type Controls interface {
	GetWallet() Wallet
}

type Subscriptions interface{}

type Wallet interface{}
