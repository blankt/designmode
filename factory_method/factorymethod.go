package factory_method

import "fmt"

/**
简单工厂中,如果要增加一种产品就需要修改工厂方法。如果增加过多就需要频繁修改该方法，工厂方法模式则不用.
增加工厂和对应产品，并从工厂中提供新增产品的获取。
优点：
	1.遵守开闭原则，修改时只需添加对应工厂类和产品，无需修改原有工厂类。
	2.使用了依赖倒置原则
缺点：
	1.每增加一个产品就需要增加一个具体产品结构体和创建该产品结构体，会导致结构体的个数增加非常快。
*/

type BunShop interface {
	generate(name string) Bun
}

type Bun interface {
	create()
}

type SCCabbageBun struct {
}

func (b SCCabbageBun) create() {
	fmt.Println("新鲜出炉的四川白菜包")
}

type SCPotatoBun struct {
}

func (b SCPotatoBun) create() {
	fmt.Println("新鲜出炉的四川土豆包")
}

type SCBun struct {
}

func (s SCBun) generate(name string) Bun {
	switch name {
	case "SCCabbageBun":
		return &SCCabbageBun{}
	case "SCPotatoBun":
		return &SCPotatoBun{}
	default:
		return nil
	}
}

type DBCabbageBun struct{}

func (b DBCabbageBun) create() {
	fmt.Println("新鲜出炉的东北热辣白菜包")
}

type DBPotatoBun struct{}

func (b DBPotatoBun) create() {
	fmt.Println("新鲜出炉的东北大酱土豆包")
}

type DBBun struct {
}

func (s DBBun) generate(name string) Bun {
	switch name {
	case "DBCabbageBun":
		return &DBCabbageBun{}
	case "DBPotatoBun":
		return &DBPotatoBun{}
	default:
		return nil
	}
}
