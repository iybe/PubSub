package pubsub

import (
	"fmt"
	"testing"
)

func Test_pub_Register(t *testing.T) {
	pub := NewPub()
	sub := NewSub()

	pub.Register(sub)

	if pub.subs.Size() != 1 {
		fmt.Printf("Expected: %d , actual: %d", 1, pub.subs.Size())
		t.Fail()
	}
}

func Test_pub_Deregister(t *testing.T) {
	pub := NewPub()
	sub := NewSub()

	pub.Register(sub)
	pub.Deregister(sub)

	if pub.subs.Size() != 0 {
		fmt.Printf("Expected: %d , actual: %d", 0, pub.subs.Size())
		t.Fail()
	}
}

func Test_pub_NotifyAll(t *testing.T) {
	pub := NewPub()
	sub := NewSub()
	expected := "test"

	pub.Register(sub)
	pub.NotifyAll(expected)

	actual := <-sub.Chan
	if actual != expected {
		fmt.Printf("Expected: %s , actual: %s", expected, actual)
		t.Fail()
	}
}
