package observer

import "testing"

func TestTemperaturer(t *testing.T) {
	weatherSubject := NewWeatherSubject()
	temperaturerObserver := new(TemperaturerObserver)
	pressureObserver := new(PressureObserver)
	weatherSubject.Attach(temperaturerObserver)
	weatherSubject.Attach(pressureObserver)
	weatherSubject.SetValues()

	weatherSubject.Remove(pressureObserver)

	weatherSubject.SetValues()
}
