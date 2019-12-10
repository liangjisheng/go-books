package smq

import "sync"

// GSmq ...
var GSmq *Smq

func init() {
	GSmq = NewSmq()
}

// Smq ...
type Smq struct {
	Brokers map[string]*Broker
	Mu      sync.RWMutex
}

// NewSmq ...
func NewSmq() *Smq {
	return &Smq{
		Brokers: make(map[string]*Broker),
	}
}

// Get ...
func (mq *Smq) Get(name string) *Broker {
	mq.Mu.RLock()
	if _, ok := mq.Brokers[name]; ok {
		defer mq.Mu.RUnlock()
		return mq.Brokers[name]
	}

	mq.Mu.RUnlock()
	mq.Mu.Lock()
	defer mq.Mu.Unlock()
	mq.Brokers[name] = NewStartedBroker(name, 1)
	return mq.Brokers[name]
}

// Close ...
func (mq *Smq) Close() {
	mq.Mu.RLock()
	mq.Mu.RUnlock()
	for _, v := range mq.Brokers {
		v.Stop()
	}
}
