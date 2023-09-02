package pubsub

import (
	"testing"
)

func TestManager_Register(t *testing.T) {
	m := NewManager()
	sub := NewSub()

	m.Register("test", sub)

	pub, ok := m.pubs.Load("test")
	if !ok {
		t.Fail()
	}

	if pub.(*Pub).subs.Size() != 1 {
		t.Fail()
	}
}

func TestManager_Deregister(t *testing.T) {
	m := NewManager()
	sub := NewSub()

	m.Register("test", sub)
	m.Deregister("test", sub)

	_, ok := m.pubs.Load("test")
	if ok {
		t.Fail()
	}
}

func TestManager_SendMsg(t *testing.T) {
	m := NewManager()
	sub := NewSub()

	m.Register("test", sub)
	m.SendMsg("test", "test")

	msg := <-sub.Chan

	if msg != "test" {
		t.Fail()
	}
}