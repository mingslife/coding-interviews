// Package question5
// 面试题 5：替换空格
package question5

import "strings"

// 题目：请实现一个函数，把字符串中的每个空格替换成“%20”。例如，输入“We are happy.”，则输
// 出“We%20are%20happy.”。

func ReplaceSpace(str string) string {
	builder := strings.Builder{}
	for _, r := range []rune(str) {
		if r == ' ' {
			builder.WriteString(`%20`)
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
