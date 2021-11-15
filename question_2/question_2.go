// Package question_2
// 面试题 2：实现 Singleton 模式
package question_2

// 题目：设计一个类，我们只能生成该类的一个实例。

type Singleton struct{}

var instance *Singleton

func GetInstance() *Singleton {
	return instance
}

func init() {
	instance = &Singleton{}
}
