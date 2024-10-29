package token

import (
	"testing"
)

// go test -v .

func TestNewJWT(t *testing.T) {
	user := NewUser(
		WithUserID(333),
		WithEmail(`liujy.willow@gmail.com`),
		WithOpenId(`111-222-333`),
	)

	engine := NewEngine()
	t.Logf("engine: %+v \r\n", *engine)
	j := NewJWT(user, engine)

	access, refresh, err := j.Generate()
	if err != nil {
		t.Fatalf("generate error: %+v", err)
	}

	t.Logf("access: %s \r\n", access)
	t.Logf("refresh: %s \r\n", refresh)

	Verify(j, access)
}
