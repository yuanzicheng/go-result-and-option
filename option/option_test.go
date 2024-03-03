package option

import "testing"

func TestIsOk(t *testing.T) {
	var x string = "any string"
	o := New(&x)
	if !o.IsSome() {
		t.Error("IsSome failed")
	}
}
