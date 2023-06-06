package main

import "fmt"

func main() {

	tv1 := &SamsungTv{
		currentChannel: 13,
		currentVolume:  30,
		tvOn:           true,
	}

	tv2 := &SonyTv{
		vol:     12,
		channel: 10,
		isOn:    false,
	}

	// Because the Sony Tv implements the "television" interface, we don't need an adapter
	tv2.turnOn()
	tv2.volumeUp()
	tv2.channelUp()
	tv2.goToChannel(70)
	tv2.turnOff()

	fmt.Println("--------------------")

	//  We need to create a SamsungTV adapter for the SamsungTV class, however
	// because it has an interface that's different from the one we want to use

	ssAdapter := &samsungAdapter{
		sstv: tv1,
	}

	ssAdapter.turnOn()
	ssAdapter.volumeUp()
	ssAdapter.channelUp()
	ssAdapter.goToChannel(70)
	ssAdapter.turnOff()

}
