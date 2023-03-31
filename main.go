package main

import "context"

func main() {
	as := PubSubMessage{Data: nil}
	ElPorvenirSiigo(context.Background(), as)
}
