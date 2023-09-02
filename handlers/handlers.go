package handlers

import "pubsub/pubsub"

type Handlers struct {
	Manager *pubsub.Manager
}

func NewHandlers() *Handlers {
	return &Handlers{
		Manager: pubsub.NewManager(),
	}
}
