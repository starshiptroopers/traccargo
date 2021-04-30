# traccargo
Traccargo is the Golang library for fetching devices information and device position updates from Traccar opensource GPS tracking system

It supports only requests for devices and device positions, also it supports live updates for devices, poistions and events from traccar websocket endpoint. 

# Usage example

```
var traccarURL = ""
var traccarToken = ""
var testDeviceID int64 = 0

func main() {
	trc, err := traccargo.NewTraccar(traccarURL, traccarToken)
	if err != nil {
		panic(err)
	}
  
  //write debug messages to the console
	trc.LogCommunicationWriter = os.Stdout
	trc.LogWriter = os.Stdout

  //get device position
	position, err := trc.Position(testDeviceID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Device %d: time %v, Position [%.6f %.6f]\n", position.DeviceID, position.FixTime, position.Latitude, position.Longitude)

  //subscribe to live traccar updates
	err = trc.SubscribeUpdates(func(m *traccargo.WsMessage) {
		fmt.Printf("new message received\n")
	})

	if err != nil {
		panic(err)
	}

  //just a rest
	time.Sleep(time.Second * 100)

	trc.Close()
}
```
