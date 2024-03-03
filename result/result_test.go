package result

import "testing"

func TestIsOk(t *testing.T) {
	var x int = 12345
	r := New(&x, nil)
	if !r.IsOk() {
		t.Error("IsOk failed")
	}
}
