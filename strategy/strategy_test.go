package strategy

import "testing"

func TestFlyWithWings(t *testing.T) {
	duck := NewDuck("MallDuck", &FlyWithWings{})
	duck.PerformFly()
}

func TestFlyNoWay(t *testing.T) {
	duck := NewDuck("FakeDuck", &FlyNoWay{})
	duck.PerformFly()
}
