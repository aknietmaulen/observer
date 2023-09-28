package main

type Subject interface {
	addToList(observer Observer)
	removeFromList(observer Observer)
	notifyAll()
}
