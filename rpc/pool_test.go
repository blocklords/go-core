package rpc

import (
	"errors"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	pool := NewPool([]string{"1", "2", "3", "4", "5"})
loop:

	curr, _ := pool.Curr()
	t.Logf("curr: %+v", curr)

	node, err := pool.Select()

	if errors.Is(err, ErrorAllLocked) {
		t.Fatalf("lock: %+v", err)
		return
	}

	if err != nil {
		goto loop
	}

	// 制造随即错误
	if node.Url() != "3" {
		t.Logf(">>>>>> 出错, 重新选择")

		p, _ := pool.pool.Load(kPool)
		t.Logf("pool: %+v", p.([]Node))
		time.Sleep(5 * time.Second)
		goto loop
	}

	t.Logf("node: %+v err : %+v", node, err)
}
