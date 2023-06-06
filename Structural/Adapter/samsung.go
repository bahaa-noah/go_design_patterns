package main

import "fmt"

type SamsungTv struct {
	currentChannel int
	currentVolume  int
	tvOn           bool
}

func (tv *SamsungTv) getVolume() int {
	fmt.Println("Samsung TV volume is", tv.currentVolume)
	return tv.currentVolume
}

func (tv *SamsungTv) setVolume(vol int) {
	tv.currentVolume = vol
	fmt.Println("Setting Samsung TV volume to", vol)
}

func (tv *SamsungTv) getChannel() int {
	fmt.Println("Samsung TV channel is", tv.currentChannel)
	return tv.currentChannel
}

func (tv *SamsungTv) setChannel(ch int) {
	fmt.Println("Setting Samsung TV channel to", ch)
	tv.currentChannel = ch
}

func (tv *SamsungTv) setOnState(tvOn bool) {
	if tvOn == true {
		fmt.Println("Samsung TV is on")
	} else {
		fmt.Println("Samsung TV is off")
	}
	tv.tvOn = tvOn
}
