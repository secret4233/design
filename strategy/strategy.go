package strategy

import "fmt"

type FlyBehavior interface {
	Fly()
}

type Duck struct {
	Name        string
	flyBehavior FlyBehavior
}

func (d *Duck) Display() {
	fmt.Printf("I'm %v", d.Name)
}

func (d *Duck) PerformFly() {
	d.flyBehavior.Fly()
}

type FlyWithWings struct{}

func (*FlyWithWings) Fly() {
	fmt.Println("I'm fly with wings.")
}

type FlyNoWay struct{}

func (*FlyNoWay) Fly() {
	fmt.Println("I can't fly.")
}

func NewDuck(name string, fly FlyBehavior) *Duck {
	return &Duck{
		Name:        name,
		flyBehavior: fly,
	}
}
