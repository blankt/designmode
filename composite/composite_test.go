package composite

import (
	"fmt"
	"testing"
	"time"
)

type depart struct {
	component
	childes []Component
}

func NewDepart(name string) *depart {
	d := &depart{
		childes: make([]Component, 0),
	}
	d.SetName(name)
	return d
}

func (o *depart) AddChild(child Component) {
	child.SetParent(o)
	o.childes = append(o.childes, child)
}

func (o *depart) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, o.Name())
	pre += " "
	for _, child := range o.childes {
		child.Print(pre)
	}
}

type user struct {
	component
}

func NewUser(name string) *user {
	u := &user{}
	u.SetName(name)
	return u
}

func (o *user) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, o.Name())
}

func TestNewComponent(t *testing.T) {
	//zx := NewDepart("zx")
	//tlf := NewUser("tlf")
	//pxy := NewUser("pxy")
	//zx.AddChild(tlf)
	//zx.AddChild(pxy)
	//zx.Print("")
	key := fmt.Sprintf("%x", time.Now().UnixNano())
	fmt.Println(key)
}
