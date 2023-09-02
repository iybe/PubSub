package slicesync

import "sync"

type SliceSYNC struct {
	Data  []interface{}
	Mutex sync.Mutex
}

func NewSliceSYNC() *SliceSYNC {
	return &SliceSYNC{
		Data: make([]interface{}, 0),
	}
}

func (s *SliceSYNC) Add(v interface{}) {
	s.Mutex.Lock()
	s.Data = append(s.Data, v)
	s.Mutex.Unlock()
}

func (s *SliceSYNC) Size() int {
	return len(s.Data)
}

func (s *SliceSYNC) Remove(v interface{}) {
	s.Mutex.Lock()
	size := len(s.Data)
	for i := 0; i < size; i++ {
		if s.Data[i] == v {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			break
		}
	}
	s.Mutex.Unlock()
}
