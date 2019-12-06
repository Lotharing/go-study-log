package base

import (
	"fmt"
	"testing"
)

/**
指导原则
接口实质上在底层用两个字段表示：
1.一个指向某些特定类型信息的指针。您可以将其视为"type"。  *int
2.数据指针。如果存储的数据是指针，则直接存储（指针内存）。如果存储的数据是一个值，则存储指向该值的指针 （取值符  & 拿到指针在存储）。
*/

//使用值接收器的方法既可以通过值调用，也可以通过指针调用。
type S struct {
	data string
}

var name string

func (s S) Read() string {
	return s.data
}
func (s *S) Write(str string) {
	s.data = str
}

//函数名称前面的括号是Go定义这些函数将在其上运行的对象的方式  也就是说函数名前边的()证明给s 构造体下可以调用此方法
func TestExample(t *testing.T) {
	//够燥key 1 value A的map
	sVals := map[int]S{1: {"A"}}
	// 你只能通过值调用 Read
	sVals[1].Read()
	// 这不能编译通过：
	//  sVals[1].Write("test")
	sPtrs := map[int]*S{1: {"A"}}
	// 通过指针既可以调用 Read，也可以调用 Write 方法
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}

//同样，即使该方法具有值接收器，也可以通过指针来满足接口。
type F interface {
	f()
}

type S1 struct{}

//接口实现    值
func (s S1) f() {}

type S2 struct{}

//接口实现	指针
func (s *S2) f() {}

func TestA(t *testing.T) {
	s1Val := S1{}
	s1Ptr := &S1{}
	//s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	//i = s2Val		因为 s2Val 是一个值，*** 而 S2 的 f 方法中没有使用值接收器 (括号前是什么就是谁的方法)
	i = s2Ptr
	fmt.Println(i)
}
