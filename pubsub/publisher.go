package pubsub

import slicesync "pubsub/utils/sliceSync"

type Publisher interface {
	Register(subscriber Subscriber)
	Deregister(subscriber Subscriber)
	NotifyAll()
}

type Pub struct {
	subs *slicesync.SliceSYNC
}

func NewPub() *Pub {
	return &Pub{
		subs: slicesync.NewSliceSYNC(),
	}
}

func (p *Pub) Register(subscriber Subscriber) {
	p.subs.Add(subscriber)
}

func (p *Pub) Deregister(subscriber Subscriber) {
	p.subs.Remove(subscriber)
}

func (p *Pub) NotifyAll(msg interface{}) {
	p.subs.Mutex.Lock()
	for _, sub := range p.subs.Data {
		sub.(Subscriber).Update(msg)
	}
	p.subs.Mutex.Unlock()
}
