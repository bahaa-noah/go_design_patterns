package main

func main() {
	// Construct two DataListener observers and
	// give each one a name
	listener1 := DataListener{
		Name: "Listener 1",
	}
	listener2 := DataListener{
		Name: "Listener 2",
	}

	// Create the DataSubject that the listeners will observe
	subj := &DataSubject{}
	// Register each listener with the DataSubject
	subj.registerObserver(listener1)
	subj.registerObserver(listener2)

	// Change the data in the DataSubject - this will cause the
	// onUpdate method of each listener to be called
	subj.ChangeItem("listener 3")
	// unregister one of the observers
	subj.unregisterObserver(listener1)
	subj.ChangeItem("listener 4")
}
