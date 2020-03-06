package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer(RandomBalancer, &RandomBalance{})
}

// RandomBalance 随机
type RandomBalance struct {
}

// DoBalance ...
func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instence")
		return
	}

	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]

	return
}
