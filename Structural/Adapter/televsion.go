package main

// Television interface to used with both TV types
type television interface {
	volumeUp() int
	volumeDown() int
	channelUp() int
	channelDown() int
	turnUp()
	turnDown()
	goToChannel(ch int)
}
