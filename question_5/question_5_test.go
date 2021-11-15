package question_5

import "testing"

func TestReplaceSpace(t *testing.T) {
	if ReplaceSpace(`We are happy.`) != `We%20are%20happy.` {
		t.Fail()
	}
}
