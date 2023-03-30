package decorator

import "fmt"

type Person interface {
	Cost() int
	Show()
}

type Peter struct {
}

func (Peter) Cost() int {
	return 0
}

func (Peter) Show() {
	fmt.Println("赤裸裸的peter")
}

type clothDecorator struct {
	person Person
}

func (*clothDecorator) Cost() int {
	return 0
}

func (*clothDecorator) Show() {
}

type Jack struct {
	clothDecorator
}

func (j Jack) Cost() int {
	return j.person.Cost() + 20
}

func (j Jack) Show() {
	j.person.Show()
	fmt.Println("peter穿上了jack")
}

type Hat struct {
	clothDecorator
}

func (h Hat) Cost() int {
	return h.person.Cost() + 10
}

func (h Hat) Show() {
	fmt.Println("peter戴上了帽子")
}
