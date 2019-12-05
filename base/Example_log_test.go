package base

import (
	"fmt"
	"testing"
)

func TestFeiBoNaQie(t *testing.T) {
	for i:=0; i< 10 ;i++  {
		fmt.Print(fibonaqie(i),"  ")
	}
}
// 递归
func fibonaqie(n int) (result int)  {
	if n==0 || n==1{
		return 1
	}else {
		return fibonaqie(n-1)+fibonaqie(n-2)
	}
}

//全局变量
var a int
var b int
//非递归
func TestFibList(t *testing.T)  {
	a:=1
	b:=1
	fmt.Print(a)
	for i:=0; i < 5; i++ {
		fmt.Print(" ",b)
		temp:=a
		a=b
		b=temp+a
	}
	fmt.Println()
}