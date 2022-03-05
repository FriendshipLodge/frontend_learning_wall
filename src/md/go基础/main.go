package main

import (
	"errors"
	"fmt"
	"time"
)

/*
go 优点
开源
编译型语言  运行告诉
语法简洁
并行处理封装
内存管理 、数组安全
*/

var ( //   因式分解写法用于声明全局变量
	a int
	b bool
)

//type 用法

//定义结构体
type name struct {
	n1 string
	n2 string
	n3 string
}

func (n *name)p()(){               //  struct method  func (结构体实例，结构体类型)  函数名(参数列表) (返回值列表)
	fmt.Println(n.n3)
}

func p1(n *name){               //  普通函数
	fmt.Println("1234")
}


//定义接口 接口命名经常用er结尾
//当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
type namer interface {
	Read()
	Write()
}

type tester interface {

}

//类型定义
type handle func(str string)

func exec(f handle) {
	f("hello")
}


//error 简单用法
func test1(n int )(int,error){
	if(n==1){
		return n,errors.New("出错了兄弟")
	}
	return n,nil

}


//简单interface使用方法

type phone interface {
	call(str string)
}

type  iphone struct{

}

type nokia struct {

}

func (n nokia)call(str string){
	fmt.Println("nokia",str)
}

func (i iphone)call(str string){
	fmt.Println("iphone",str)
}

func interfacetest()  {
	var  p1 phone=new(iphone)
	p1.call("123")

	p1=new(nokia)
	p1.call("666")

}

//goroutine 简单用法 goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。使用时 使用 go来开启goroutine

func pp1(str string)  {
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(str)

	}


}
func goroutineTest(){
	go pp1("hello")
	pp1("world")
}

func main() { // {不能放在单独的行上


	goroutineTest()
	//interfacetest()

	_,e:=test1(1)
	if e!=nil{
		fmt.Println(e)
	}


	//数组
	const aa int = 3
	var numbers = [3]int{}
	var nums2 = [3]int{}
	var nums3 = [aa]int{}
	var nums4 = [...]int{}
	var nums5 = [5]int{1, 2, 3, 4, 5}
	var size int = len(nums4)
	numbers[1] = 1

	fmt.Println(numbers[1], nums2[0], nums3, nums4, size) // num3 输出整个数组

	//slice 可以理解为动态数组
	var slice = nums5[1:3] // index左闭右开
	s1 := []int{1, 2, 3}

	s2 := nums5[:] //空白表示头或者尾 s2:=nums5[0:]  0到尾 s2:=nums[:2] 头到1

	var capslice int = cap(slice) //cap=数组长度-左index 的绝对值
	slice = slice[:capslice]      //由 2 3  拓展为 2345
	slice = append(slice, 1, 2)   //拓展元素 1 2 一旦扩展操作超出了被操作的切片值的容量，那么该切片的底层数组就会被替换
	fmt.Println(slice, capslice, s1, s2)

	var s3 = []int{0, 0, 0, 0, 0}
	var s4 = []int{1, 2, 3}
	copy(s3, s4) //从 第一个元素开始替换
	fmt.Println(s3)

	//变量声明初始化
	var c int = 10
	var d = 20
	ee := "weipai " // 只能在函数体内使用 自动声明和初始化

	//定义常量
	const LENGTH int = 10
	//枚举 遇到一个const itoa被置为0 每新增一行常量声明将使iota+1
	const (
		UnKnown = 0        //iota=0
		Female  = 2 * iota //=1
		Male    = iota     //=2
	)

	//局部变量声明之后必须被使用 不然会报错。
	fmt.Println("hello world", c, d, ee, Male)

	//使用这种方法增加可读性，如果handle 有五个参数，exec只需要传一个handle类型就ok
	var p handle = func(str string) {
		str += "1"
		fmt.Println("123", str)
	}
	exec(p)

	//匿名函数作为参数
	exec(func(str string) {
		fmt.Println("312", str)
	})

	// switch 用法
	marks := 90
	switch marks {
	case 90:
		fmt.Println("2")
		// fallthrough 不会判断下一条语句是否为真 直接执行
	case 70:
		fmt.Println("70")
		//fallthrough 同上
	case 80:
		fmt.Println("成绩是90")

	default:
		fmt.Println("weizhi")
	}

	// switch 被用于 tpyewitch 来判断某个interface变量中实际 储存的变量类型
	// 使用type进行类型查询时，只能在switch中使用，且使用类型查询的变量类型必须是interface{}
	var x interface{}
	switch i := x.(type) {
	case nil: //nil 空类型
		fmt.Println("x 的类型是 %T", i) //%T 使用go语法 输出变量的类型
	}

	var str1 string = "\\str1" //输出 \str1
	var str2 string = `\\str2` //输出 \\str2  所见即所得
	str1 += str2
	fmt.Println(str1)



	//interface 类型
	var nnn name=name{n1:"123",n2:"123",n3:"123"}
	nnn.p()
	//p(&nnn)


	p1(&nnn)
	//var nani *name =&nnn
	///????????????????????????????????????????????p(nani)
	var nn interface{}=nnn
	fmt.Printf("%+v\n",nnn)  // printf 根据format打印值 %+v 打印形式如 n1:123
	fmt.Printf("%+v\n",nn)   //%v打印结构体 或interface里的值




	//select chan make 三种for循环

	//range Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	//在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值。


	arr := []int{2,3,4,5}
	sum := 0
	// range 可以返回数组or切片中元素的下标和元素的值，此时如果不需要使用下标时，可以用_表示省略
	for index, num := range arr {
		fmt.Println(index, num)
		sum += num
	}

	//var nums6 = [5]int{1, 2, 3, 4, 5}
	//for index,num:=range nums6{
	//	fmt.Println(index,num)
	//}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	kvs["1"]="1"
	for v := range kvs {
		fmt.Printf("%s -> %s\n",  v)
	}


	//while
	aaa:=1
	bbb:=3
	for  aaa<bbb {
		fmt.Println(aaa)
		aaa++
	}


	/*一些特性
	1： 嵌入类型
	可以使内部类型提升到外部类型中。被提升的内部类型标识符就好像直接声明在外部类型中。可以直接通过外部类型访问。
	当外部类型重新声明时，内部相同标识符不会被提升。
	e.g.：

	// user 在程序里定义一个用户类型
	type user struct {
	    name string
	    email string
	}
	// notify 实现了一个可以通过user 类型值的指针
	// 调用的方法
	func (u *user) notify() {
	    fmt.Printf("Sending user email to %s<%s>\n",
	    u.name,
	    u.email)
	}
	// admin 代表一个拥有权限的管理员用户
	type admin struct {
	    user // 嵌入类型
	    level string
	}

	定义一个ad 为admin类型
	// 我们可以直接访问内部类型的方法
	ad.user.notify()
	// 内部类型的方法也被提升到外部类型
	ad.notify()

	//-------------------------------------------------------//
	// 通过user 类型值的指针
	// 调用的方法
	func (u *user) notify() {
		fmt.Printf("Sending user email to %s<%s>\n",
			u.name,
			u.email)
	}



	// 通过 admin 类型值的指针
	// 调用的方法
	func (a *admin) notify() {
		fmt.Printf("Sending admin email to %s<%s>\n",
			a.name,
			a.email)
	}

	此时可分别调用
	// 我们可以直接访问内部类型的方法
	ad.user.notify()
	// 内部类型的方法没有被提升
	ad.notify()

	2：select


	每个 case 都必须是一个通信
	所有 channel 表达式都会被求值
	所有被发送的表达式都会被求值
	如果任意某个通信可以进行，它就执行，其他被忽略。
	如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
	否则：
	如果有 default 子句，则执行该语句。
	如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

	3：defer关键字
	defer 后面的函数值和函数参数会被计算 但是不会被调用
	实际上调用要等外围函数调用完之后调用

	多个defer出现时，defer函数出现的顺序与执行顺序相反

	e.g. ：
	f,err := os.Open(fileName)
	if err!=nil {
	return
	}
	defer f.Close()
	*/




}


