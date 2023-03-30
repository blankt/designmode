package strategy

import "fmt"

/*
	策略模式:定义一系列的算法，把它们一个个封装起来，并且使它们可以互相替换。
	应用实例：
	1、主题的更换，每个主题都是一种策略
	2、旅行的出游方式，选择骑自行车、坐汽车，每一种旅行方式都是一个策略。
*/

type Vehicle interface {
	Run()
}

type Car struct {
}

func (c Car) Run() {
	fmt.Println("汽车开始行驶")
}

type Airplane struct {
}

func (air Airplane) Run() {
	fmt.Println("飞机开始起飞")
}

type Traveler struct {
	vehicle Vehicle
}

func (t *Traveler) SetTraveler(vehicle Vehicle) {
	t.vehicle = vehicle
}

func (t *Traveler) Travel() {
	t.vehicle.Run()
}
