package result

import "testing"

func TestResult(t *testing.T) {
	var x int = 12345
	r := New(&x, nil)
	if !r.IsOk() {
		t.Error("IsOk failed")
	}

	v := r.Unwrap()
	if v != x {
		t.Error("Unwrap failed")
	}
}
