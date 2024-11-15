package rpc

import (
	"errors"
	"github.com/blocklords/go-core/fn"
	"sync"
)

const (
	defaultWait = 900
	kPool       = `pool`
	kIndex      = `index`
)

var (
	ErrorAllLocked = errors.New(`all node are locked`)
)

type (
	PoolName string

	Node struct {
		url   string // rcp
		wait  int64  // rpc 出错的情况下锁定时间
		index bool   // 当前是否正在使用
	}

	Pool struct {
		pool *sync.Map
	}
)

func (n *Node) Url() string {
	return n.url
}

func (n *Node) Wait() int64 {
	return n.wait
}

func NewPool(rpcs []string) *Pool {
	if len(rpcs) < 1 {
		panic(`bad rpcs`)
	}
	ns := make([]Node, 0)
	for _, rpc := range rpcs {
		ns = append(ns, Node{
			url:   rpc,
			wait:  0,
			index: false,
		})
	}

	ns[0].index = true

	store := &sync.Map{}
	store.Store(kPool, ns)
	store.Store(kIndex, 0)

	return &Pool{pool: store}
}

func (p *Pool) Curr() (*Node, error) {
	index, e := p.pool.Load(kIndex)
	if !e {
		return nil, errors.New(`bad index`)
	}

	pool, _ := p.pool.Load(kPool)

	return &(pool.([]Node)[index.(int)]), nil
}

func (p *Pool) Select() (*Node, error) {
	now := fn.Now()
	var n *Node

	pool, _ := p.pool.Load(kPool)

	ps := pool.([]Node)

	index := 0
	for i, node := range ps {
		if node.wait > now {
			continue
		}

		if node.index {
			ps[i].wait = now + defaultWait
			ps[i].index = false
			continue
		}

		ps[i].index = true
		index = i
		n = &node
		break
	}

	p.pool.Store(kPool, ps)
	p.pool.Store(kIndex, index)

	if n == nil {
		return n, ErrorAllLocked
	}

	return n, nil
}
