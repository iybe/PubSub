package pubsub

type Subscriber interface {
	Update(interface{})
}

type Sub struct {
	Chan chan (interface{})
}

func NewSub() *Sub {
	return &Sub{Chan: make(chan interface{}, 1)}
}

func (s *Sub) Update(msg interface{}) {
	s.Chan <- msg
}
