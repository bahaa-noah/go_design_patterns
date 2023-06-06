package main

type samsungAdapter struct {
	sstv *SamsungTv
}

func (ss *samsungAdapter) turnOn() {
	ss.sstv.setOnState(true)
}

func (ss *samsungAdapter) turnOff() {
	ss.sstv.setOnState(false)
}

func (ss *samsungAdapter) volumeUp() int {
	vol := ss.sstv.getVolume() + 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *samsungAdapter) volumeDown() int {
	vol := ss.sstv.getVolume() - 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *samsungAdapter) channelUp() int {
	channel := ss.sstv.getChannel() + 1
	ss.sstv.setChannel(channel)
	return channel
}

func (ss *samsungAdapter) channelDown() int {
	channel := ss.sstv.getChannel() - 1
	ss.sstv.setChannel(channel)
	return channel
}

func (ss *samsungAdapter) goToChannel(ch int) {
	ss.sstv.setChannel(ch)
}
