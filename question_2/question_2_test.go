package question_2

import "testing"

func TestGetInstance(t *testing.T) {
	if GetInstance() == nil {
		t.Fail()
	}
}
