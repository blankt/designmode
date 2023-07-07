package facade

import "log"

/**
Facade: 观察者模式-结构型
	将子系统的接口进行组装，提供给外部一个简单的接口。
	例如目前代码中需要调用子系统中的众多复杂类及其方法，这时候在程序中加入调用就会使当前代码与子系统高度耦合，这时候创建一个外观类封装调用过程并提供一个简单的接口，就会使当前代码与子系统中具体实现解耦，如果子系统有变动，
修改外观类即可。
	优点：让自己的代码独立于复杂子系统
	缺点：外观对象可能成为上帝对象（了解过多或者负责过多的对象-对象中涉及的模块、操作很多）

*/

type AModule interface {
	TestA()
}

type aModuleImpl struct {
}

func (a aModuleImpl) TestA() {
	log.Println("a module running")
}

func NewAModule() AModule {
	return aModuleImpl{}
}

type BModule interface {
	TestB()
}

type bModuleImpl struct {
}

func (b bModuleImpl) TestB() {
	log.Println("b module running")
}

func NewBModule() BModule {
	return bModuleImpl{}
}

type Api interface {
	Test()
}

type apiImpl struct {
	a AModule
	b BModule
}

func (a apiImpl) Test() {
	a.a.TestA()
	a.b.TestB()
}

func NewApi() Api {
	return apiImpl{
		a: NewAModule(),
		b: NewBModule(),
	}
}
