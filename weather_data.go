package main

type WeatherData struct {
	observers []Observer
	message   string
}

func newWeatherData() *WeatherData {
	return &WeatherData{}
}

func (wd *WeatherData) addToList(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) removeFromList(o Observer) {
	for i, obs := range wd.observers {
		if obs == o {
			wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
			break
		}
	}
}

func (wd *WeatherData) notifyAll() {
	for _, o := range wd.observers {
		o.update(wd.message)
	}
}

func (wd *WeatherData) setMessage(message string) {
	wd.message = message
	wd.notifyAll()
}
