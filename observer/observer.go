package observer

import "fmt"

type Subject interface {
	Attach(Observer)
	Remove(Observer)
	Notify()
}

type WeatherSubject struct {
	observers   []Observer
	temperature float64
}

var _ Subject = (*WeatherSubject)(nil)

func NewSubject() *WeatherSubject {
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

type Observer interface {
	Update(Subject)
}

var _ Observer = (*temperaturerObserver)(nil)

type temperaturerObserver struct {
}

func (w *temperaturerObserver) Update(s Subject) {
	fmt.Printf("now temperature is:%v", s.(*WeatherSubject).temperature)
}
