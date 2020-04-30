package singleflight

import "sync"

// Call 代表正在进行中 或已经结束的请求 使用 sync.WaitGroup 锁避免重入
// is an in-flight or completed Do call
type Call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group 是 singleflight 的主数据结构 管理不同 key 的请求(call)
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*Call // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*Call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(Call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
