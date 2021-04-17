package observer

import (
	"fmt"
	"math/rand"
)

// 主题
type Subject interface {
	Attach(Observer)
	Remove(Observer)
	Notify()
}

type WeatherSubject struct {
	observers   []Observer
	temperature float64
	pressure    float64
}

var _ Subject = (*WeatherSubject)(nil)

func NewWeatherSubject() *WeatherSubject {
	return &WeatherSubject{
		observers: make([]Observer, 0),
	}
}

func (s *WeatherSubject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *WeatherSubject) Remove(o Observer) {
	var index int = -1
	for i, nowObserver := range s.observers {
		if o == nowObserver {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	s.observers = append(s.observers[:index], s.observers[index+1:]...)
}

func (s *WeatherSubject) Notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *WeatherSubject) SetValues() {
	s.pressure = rand.Float64()
	s.temperature = rand.Float64()
	s.Notify()
}

// 观察者
type Observer interface {
	Update(Subject)
}

var _ Observer = (*TemperaturerObserver)(nil)

type TemperaturerObserver struct {
}

func (w *TemperaturerObserver) Update(s Subject) {
	fmt.Printf("now temperature is:%v\n", s.(*WeatherSubject).temperature)
}

type PressureObserver struct {
}

var _ Observer = (*PressureObserver)(nil)

func (p *PressureObserver) Update(s Subject) {
	fmt.Printf("now pressure is:%v\n", s.(*WeatherSubject).pressure)
}
