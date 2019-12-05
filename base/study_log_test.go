package base
//test 不是包命，就是以test结尾  文件名以xxx_test.go命名
import (
	"fmt"
	"math"
	"testing"	//测试包
)
//常量定义/初始化

/**
	1.函数内定义的变量称为局部变量
	2.函数外定义的变量称为全局变量
	3.函数定义中的变量称为形式参数， 形参会作为函数的局部变量来使用
	注： go中全局变量和局部变量名称可以相同，但是函数内部的局部变量会被优先考虑
	pointer 的 初始化为nil		int  float32 初始化 0
 */

const(
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

//测试方法名称必须Test开头
/**
	iota 基本使用
 */
func TestConstantTryOne(t *testing.T)  {
	fmt.Println(Monday,Tuesday,Wednesday)
}
//多参复制
func TestExchange(t *testing.T)  {
	a:=1
	b:=2
	//temp:=a
	//a=b
	//b=temp
	a,b = b,a
	fmt.Println(a," ",b)
}
/**
	位与运算
 */
func TestConstantTryTwo(t *testing.T)  {
	a:=7 //0111 	二进制
	a =  a&^Readable
	a =  a&^Writable
	fmt.Println(a&Readable==Readable,a&Writable==Writable,a&Executable==Executable)
}
/**
	类型转换
 */
//别名
type MyInt int64
func TestImplicit(t *testing.T)  {
	var a int32 = 1
	var b int64
	//显式类型转换
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	fmt.Println(a,b,c)
}

/**
	预定义类型
 */
func TestReType(t *testing.T)  {
	var a = math.MaxInt64
	var b = math.MaxFloat32
	var c = math.MaxUint32
	fmt.Println(a,b,c)
}
/**
	指针类型: 不支持指针运算，有垃圾回收，string是值类型，默认初始化值是空 而不是nil , go 语言中不支持指针运算
 */
func TestPoint(t *testing.T)  {
	a:=1
	// & 为指针取址符
	aPtr := &a
	fmt.Println(a," ",aPtr)
	//采用格式化符，与C类似
	//fmt.Println("%T", a)
}
/**
	string 的默认值为 空  长度为0
 */
func TestString(t *testing.T)  {
	var a string
	fmt.Println("*"+a+"*")
	fmt.Println(len(a))
}

/****
	+ - * / %  ++  --    ==  ！=  >   <   >=   <=
	运算符号  go中没有 ++前置  例如 ++i
	== 比较数组，相同维数，相同元素个数的的数组可比较 || 每个元素都相同的才相等

	在java中比较的是数组的引用，在go中是直接比对数组的内容
****/

func TestCompareArray(t *testing.T)  {
	a := [...]int{1, 2 , 3 ,4}
	b := [...]int{1, 2 , 3 ,4}
	d := [...]int{1, 3 , 2 ,4}
	e := [...]int{1, 2 , 3 ,5}
	//c := [...]int{1, 2 , 3 ,4, 5}		数组长度不相同比较，直接有错
	fmt.Println(a==b)		//值与位置都相同		true
	fmt.Println(a==d)		//值相同但位置不同	false
	fmt.Println(a==e)		//值不同				false
}

/**
	逻辑运算符：
		&& 逻辑AND运算符 两边都是True 则条件为True
		|| 逻辑OR运算符 两边一个为True 则条件为True
		！ 逻辑Not运算符 如果条件为True  则逻辑NOT条件False
	位运算符 和 主流语言一直

	&^ 按位置零	案例：函数名 TestConstantTryTwo
	0 &^ 1  - 1 右边是1 ，左边无论什么都清0
	0 &^ 0		右边是0，则左边是什么就保留什么
 */

/**
	Go语言仅支持循环关键字for 但是与Java不同是不加 （）
	//for 写法
	for i:=0 ； i < 3 ； i++
	while: //循环写法
	for n < 5{
		n++
	}
	while(true):
	for{
		//死循环
	}
	//if语句		go中的if语句支持变量的赋值
	if condition {} 	不需要()  condition必须是bool条件
	if var declaration; condition {}
 */
func TestIfMultiSec(t *testing.T)  {
	if a := 1 == 1; a {
		fmt.Println("1==1")
	}
	// v 接收方法的返回值， 如果没有err 就是true 则就继续执行
	//if v,err := someFUn(); err==nil {
	//
	//}else {
	//
	//}
}

/**
		Switch语句
	1.条件表达式不限制常量或者整数  字符串也可以
	2.单个case中，可以出现多个结果选项，使用逗号分割
	3.GO语言中不需要用break退出一个case
	4.可以不设定switch之后得条件表达式，整个switch机构与多个if else得逻辑作用相同
 */

func TestSwitchMultiCase(t *testing.T)  {
	for i := 0 ; i < 5 ; i++  {
		switch i {
		case 0,2:
			fmt.Println("Even ")
		case 1,3:
			fmt.Println("Odd")
		default:
			fmt.Println("it is not 0 - 3")
		}
	}
	/**
		  if else  作用
	 */
	for i := 0 ; i < 5 ; i++  {
		switch {
		case i%2 == 0:
			fmt.Println("Even")
		case i%2 == 1:
			fmt.Println("Odd")
		default:
			fmt.Println("un know")
		}
	}
}
/**
	集合： 数组和切片
 */

func TestArray(t *testing.T)  {
	//声明并初始化为默认值0	三个位置都是默认0  对应类型的0值  string - nil
	var a [3]int;
	a[0] = 1
	fmt.Println(a[0])
	//声明同时初始化
	b := [3]int{1,2,3}
	//遍历数组 len()函数可以得到数组的长度
	for i := 0; i < len(b)  ;i++  {
		fmt.Println(b[i])
	}
	//多维数组初始化
	c := [2][2]int{{1,2},{3,4}}
	fmt.Println("=====多维数组遍历=====")
	//多维数组遍历
	for i := 0 ; i  < len(c) ; i++  {
		for j := 0 ; j  < len(c) ; j++  {
			fmt.Println(c[i][j])
		}
	}
	// ...代表自动识别初始化数组几位
	d := [...]int{1,2,3,4,5}
	fmt.Println(d[3])
	fmt.Println("=====range=====")
	//类似 foreach 的语法  使用range关键字   index 是 索引， e是值  输出[索引，值]
	for index, e := range d{
		fmt.Println(index,e)
	}
	fmt.Println("======‘_的作用’====")
	//下划线代表并不关心这个值的结果
	for _, e := range d{
		fmt.Println(e)
	}
	fmt.Println("======数组的截取====")
	//数组截取  a[开始索引（包含），结束索引（不包含）]
	fmt.Println(d[1:2])  //输出[2] type : 数组
	fmt.Println(d[1:len(d)]) //2 3 4 5
	fmt.Println(d[1:])  //2 3 4 5 左边包含
	fmt.Println(d[:3])  //1 2 3   右边不包含
	fmt.Println(d[:]) // 全部截取
}
/**
	方法调用简答练习 1.单个返回值 2.多个返回值

	小知识点： 	 值传递：值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
				 引用传递：引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

	小知识点：    ：= 是声明变量并赋值 		例如： var a = fun();  //  a := fun()	可以省略var
 */
func TestDialog(t *testing.T)  {
	var a = max(1,3)
	fmt.Println(a)
	//多个返回值函数接收调用
	var m,n = swap("lutong","niubi")
	fmt.Println(m,n)
}
func max(num1,num2 int) int  {
	var result int
	if num1 > num2 {
		result = num1
	}else {
		result = num2
	}
	return result
}
//多个返回值
func swap(x, y string) (string,string)  {
	return y,x
}

/**
	指针的简单了解，在java中指针也就是所谓的引用， 它指向的是数据在计算机中的内存地址
	在Go中，取指符 为& 放在变量前就会返回相应的的地址
	指针的声明格式：var var_name *var-type		指针名称 ， * 号用于指定这个变量为指针，  var-type指针类型

	定义的指定变量 是nil的， 其值为对应指针类型的 默认值 如 var ptr *int  空指针，默认值是 0
	判断指针的方法是： ptr != nil 就不是空指针，不是空指针就代表 指向的有数据的地址

	扩展：
		1.Go指针数组
		2.Go指向指针的指针
		3.向函数传递指针参数
 */

func TestPointExample(t *testing.T)  {
	var a int = 20
	var ip *int
	//指针指向a的地址 / 指针变量的存储地址
	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a )
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	fmt.Printf("*ip 变量的值: %d\n", *ip )
}

/**

	语言结构体：
		Go语言在数组中可以存储相同类型的数据，在struct结构体中可以存储不同类型的数据，是一种数据集合，表示一项记录，与java中的实体相应
		定义结构体类型： 可以用于变量的声明
		1. variable_name := structure_variable_type {value1, value2...valuen}
		2. variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
 */
// 创建机构体
type Books struct {
	title string
	author string
}
// struct 结构体声明案例
func TestStructExample(t *testing.T)  {
	// key - value 格式声明一个结构体
	fmt.Println(Books{
		title:  "数学",
		author: "路通",
	})
	//直接按照参数顺序声明
	fmt.Println(Books{"Go语言","路通"})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言"})

	var book1 Books
	book1.author="Go 语言"
	book1.title="路通"
	fmt.Println(book1.title)

	//结构体指针
	var ptr *Books
	ptr = &Books{
		title:  "GO",
		author: "指针作者",
	}
	printBooks(ptr)
}
//指针参数
func printBooks(Books *Books)  {
	fmt.Println(Books.author)
}
//如果想在函数里面改变结果体数据内容，需要传入指针  指针.属性 = 值

/**
	切片
	声明方式：
	var s0 []int
	s0 = append(s0,1)
	s := []int{}
	s1 := []int{1,2,3}
	//创建一个长度为2 容量为4的
	s2 := make([]int,2,4)

	//数组		也就是说数组在声明时需要指定长度，或者用 [...]自动识别长度，切片不需要，所以定义时候就看此区别
	a := [...]int{1, 2 , 3 ,4}

	切片数组区别：
		1：切片容量可伸缩，数组不能
		2：数组可以进行比较，切片不能

 */
func TestSliceInti(t *testing.T)  {
	//slice声明，不需要指定长度 		len长度 cap容量
	var s0 []int
	fmt.Println(len(s0),"",cap(s0))
	s0 = append(s0,1)
	fmt.Println(len(s0),"",cap(s0))

	s1 := []int{1,2,3,4}
	fmt.Println(len(s1),"",cap(s1))
	//初始化几个就只能访问几个  3
	s2 := make([]int,3,5)
	fmt.Println(len(s2),"",cap(s2))
	fmt.Println(s2[0],s2[1],s2[2])	//初始化三个就只能访问三个
	s2 = append(s2,3)			//长度增加1
	fmt.Println(len(s2),cap(s2))
	fmt.Println(s2[0],s2[1],s2[2],s2[3])
	fmt.Println("=====")
	temp := make([]int,2,4) //len个元素会被初始化默认为 0 未初始化元素不可访问
	temp = append(temp,1 )
	temp = append(temp,2 )
	temp = append(temp,3 )
	temp = append(temp,4 )
	temp[0] = 5
	fmt.Println(temp)
}

/**
	切片共享存储结构，自动变长
 */
func TestSliceGrowing(t *testing.T)  {
	s := []int{}
	for i:=0;i<10 ;i++  {
		s = append(s, i)
		fmt.Println(len(s),cap(s))		//当到第九个数的时候，切片长度由 8 -> 16
	}

	words := []string{"a","b","c","d","e","f","g","h","i","j","k","l"}
	word := words[3:6]  //截取三个
	fmt.Println(len(word),cap(word))		//长度就是三，从截取位置往后共有九个数据，所以cap容量数9
	fmt.Println(word)
	temp := words[5:8]
	fmt.Println(len(temp),cap(temp))
	temp[0] = "Unknow"
	fmt.Println(word)	//temp中的第一个数据，下标 6 更改为unknow  word的最后一个变unknow  数据共享空间
	fmt.Println(words)
}

/**
	map(集合) 无序 键值对
	var var map_variable map[key_data_type]value_data_type
	内建函数make:
	map_variable := make(map[key_data_type]value_data_type)

	不初始化map 就会创建一个nil的map nil的map 不能存放键值对
 */
func TestMapExample(t *testing.T)  {
	//interface{} 代表value可以是任何类型
	var countryCapitalMap map[string]interface{}
	countryCapitalMap = make(map[string]interface{})
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	for country := range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}
	// 返回 value  存在
	capital, ok := countryCapitalMap [ "Japan" ] /*如果确定是真实的,则存在,否则不存在 */
	if (ok) {
		fmt.Println("Japan 的首都是", capital)
	} else {
		fmt.Println("Japan 的首都不存在")
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	for country := range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}

}
