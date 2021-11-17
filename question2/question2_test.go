package question2

import "testing"

func TestGetInstance(t *testing.T) {
	if GetInstance() == nil {
		t.Fail()
	}
}
