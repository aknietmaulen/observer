package main

type Observer interface {
	update(message string)
	getID() string
}
