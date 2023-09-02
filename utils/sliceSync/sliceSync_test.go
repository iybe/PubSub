package slicesync

import (
	"fmt"
	"sync"
	"testing"
)

func Test_sliceSYNC_Add_multiples_goroutines(t *testing.T) {
	slice := NewSliceSYNC()
	lenExpected := 10

	wg := sync.WaitGroup{}
	for i := 0; i < lenExpected; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			slice.Add(i)
		}(i)
	}
	wg.Wait()

	lenActual := len(slice.Data)
	if lenActual != lenExpected {
		fmt.Printf("Expected: %d , actual: %d", lenExpected, lenActual)
		t.Fail()
	}
}

func Test_sliceSYNC_Size(t *testing.T) {
	slice := NewSliceSYNC()
	lenExpected := 2

	slice.Add(1)
	slice.Add("123")

	lenActual := slice.Size()
	if lenActual != lenExpected {
		fmt.Printf("Expected: %d , actual: %d", lenExpected, lenActual)
		t.Fail()
	}
}

func Test_sliceSYNC_Remove(t *testing.T) {
	slice := NewSliceSYNC()
	lenExpected := 2

	slice.Add(1)
	slice.Add(2)
	slice.Add(3)
	slice.Add(4)
	wg := sync.WaitGroup{}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			slice.Remove(i)
		}(i)
	}
	wg.Wait()

	lenActual := len(slice.Data)
	if lenActual != lenExpected {
		fmt.Printf("Expected: %d , actual: %d", lenExpected, lenActual)
		t.Fail()
	}
}
