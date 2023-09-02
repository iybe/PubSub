package pubsub

import (
	"sync"
)

type Manager struct {
	pubs sync.Map
}

func NewManager() *Manager {
	return &Manager{
		pubs: sync.Map{},
	}
}

func (m *Manager) Register(idPub string, sub Subscriber) {
	pub, ok := m.pubs.Load(idPub)
	if !ok {
		pub = NewPub()
		m.pubs.Store(idPub, pub)
	}
	pub.(*Pub).Register(sub)
}

func (m *Manager) Deregister(idPub string, sub Subscriber) {
	pub, ok := m.pubs.Load(idPub)
	if !ok {
		return
	}
	pub.(*Pub).Deregister(sub)
	if pub.(*Pub).subs.Size() == 0 {
		m.pubs.Delete(idPub)
	}
}

func (m *Manager) DeregisterSafe(idPub string, sub Subscriber) {
	pub, ok := m.pubs.Load(idPub)
	if !ok {
		return
	}
	pub.(*Pub).Deregister(sub)
	close(sub.(*Sub).Chan)
	if pub.(*Pub).subs.Size() == 0 {
		m.pubs.Delete(idPub)
	}
}

func (m *Manager) SendMsg(idPub string, msg interface{}) {
	pub, ok := m.pubs.Load(idPub)
	if !ok {
		return
	}
	pub.(*Pub).NotifyAll(msg)
}
