package pubsub

import (
	"fmt"
	"testing"
)

func TestSub_Update(t *testing.T) {
	sub := NewSub()
	expected := "test"

	sub.Update(expected)
	close(sub.Chan)

	actual := <-sub.Chan
	if actual != expected {
		fmt.Printf("Expected: %s , actual: %s", expected, actual)
		t.Fail()
	}
}
