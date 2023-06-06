package main

import "fmt"

type SonyTv struct {
	vol     int
	channel int
	isOn    bool
}

func (st *SonyTv) turnOn() {
	fmt.Println("Sony Tv is on")
	st.isOn = true
}

func (st *SonyTv) turnOff() {
	fmt.Println("Sony Tv is off")
	st.isOn = false
}

func (st *SonyTv) volumeUp() int {
	st.vol++
	fmt.Println("Increasing sony tv volume to", st.vol)
	return st.vol
}

func (st *SonyTv) volumeDown() int {
	st.vol--
	fmt.Println("Decreasing sony tv volume to", st.vol)
	return st.vol
}

func (st *SonyTv) channelUp() int {
	st.channel++
	fmt.Println("Increasing sony tv channel to", st.channel)
	return st.channel
}

func (st *SonyTv) channelDown() int {
	st.channel--
	fmt.Println("Decreasing sony tv channel to", st.channel)
	return st.channel
}

func (st *SonyTv) goToChannel(ch int) {
	st.channel = ch
	fmt.Println("Setting Sony tv channel to", st.channel)

}
