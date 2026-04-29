package rand_any

import "testing"

// go test github.com/blocklords/go-core/rand -run TestRandCode_Make -v
func TestRandCode_Make(t *testing.T) {
	rc := NewRandCode(
	//RandCodeFormat(NewFormatString(FormatStringGroup(0))),
	)
	for i := 0; i < 10; i++ {
		code := rc.Make()
		t.Logf("code: %s", code)
	}
	rc1 := NewRandCode(
		RandCodeFormat(NewFormatString(FormatStringGroup(3))),
	)
	for i := 0; i < 10; i++ {
		code := rc1.Make()
		t.Logf("code: %s", code)
	}
}
