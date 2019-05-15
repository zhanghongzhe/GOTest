package Test3

import "fmt"

func Test()  {
	//初始化
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{height: 20}
	rect4 := &Rect{10,20}

	fmt.Println(rect1,rect2,rect3)
	fmt.Println(rect4.Area())
}

//定义一个对象
type Rect struct {
	width, height int
}

//定义对象方法
func (r *Rect) Area() int {
	return r.width * r.height
}

//在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以 NewXXX 来命名，表示“构造函数”：
func NewRect(width, height int) *Rect {
	return &Rect{width, height}
}