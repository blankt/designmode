package template

import "fmt"

/*
	模板模式：定义一个操作中的算法的骨架，而将一些步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
	关键：通用步骤在抽象类中实现，关键的步骤在具体的子类中实现
	实例：
	1、做饭步骤：进厨房、开火、（做饭）、关火、结束，除了做饭这一步骤，其他步骤都是相同的，可以继承抽象类自己定义这一步骤，只修改这一步骤。
*/

type Cooker interface {
	Open()
	Fire()
	Cook()
	OutFire()
	Close()
}

type Menu struct {
}

func (Menu) Open() {
	fmt.Println("进入厨房")
}

func (Menu) Fire() {
	fmt.Println("开火")
}

func (Menu) Cook() {
}

func (Menu) OutFire() {
	fmt.Println("关火")
}

func (Menu) Close() {
	fmt.Println("结束")
}

func Cooking(cook Cooker) {
	cook.Open()
	cook.Fire()
	cook.Cook()
	cook.OutFire()
	cook.Close()
}

type Tomato struct {
	Menu
}

func (t Tomato) Cook() {
	fmt.Println("炒西红柿")
}

type Potato struct {
	Menu
}

func (t Potato) Cook() {
	fmt.Println("炒土豆")
}
