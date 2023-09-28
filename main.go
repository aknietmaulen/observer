package main

//This observer pattern is used to send notifications of Weather forecast to people's phone numbers

func main() {
	weatherData := newWeatherData()

	person1 := newPerson("87755634918")
	person2 := newPerson("87777777777")

	weatherData.addToList(person1)
	weatherData.addToList(person2)

	weatherData.setMessage("Weather forecast: It's going to snow")
}
