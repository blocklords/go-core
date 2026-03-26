package rand_any

import "testing"

// go test github.com/blocklords/go-core/rand -run TestRandCode_Make -v
func TestRandCode_Make(t *testing.T) {
	for i := 0; i < 10; i++ {
		rc := NewRandCode(
			//RandCodeFormat(NewFormatString(FormatStringGroup(0))),
		)
		code := rc.Make()
		t.Logf("code: %s", code)
	}
}
