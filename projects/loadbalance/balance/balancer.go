package balance

import "fmt"

var (
	// RandomBalancer 随机
	RandomBalancer = "random"
	// RoundRobinBalancer 轮询
	RoundRobinBalancer = "roundrobin"
	// HashBalancer hash
	HashBalancer = "hash"
)

// Balancer ...
type Balancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}

// BalanceMgr ...
type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var mgr = BalanceMgr{
	allBalancer: make(map[string]Balancer),
}

func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

// RegisterBalancer ...
func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

// DoBalance ...
func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balance", name)
		return
	}
	fmt.Printf("use %s balancer: ", name)
	inst, err = balancer.DoBalance(insts)
	return
}
