package question6

import "testing"

func TestReverseKey(t *testing.T) {
	node := &ListNode{
		Key: 1,
		Next: &ListNode{
			Key: 2,
			Next: &ListNode{
				Key: 3,
			},
		},
	}
	if ReverseKey(node) != "321" {
		t.Fail()
	}
}
