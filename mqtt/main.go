package main

import (
	"fmt"
	"net/url"
	"time"

	emitter "github.com/emitter-io/go"
)

const (
	key     = "Vv-L7ADx1qh506VMwFBFyuhmBPMiA3kj"
	channel = "sensor/"
)

func main() {
	// Create the options with default values
	o := emitter.NewClientOptions()
	o.Servers = []*url.URL{
		{
			Scheme: "tcp",
			Host:   "localhost:8081",
		},
	}

	// Set the message handler
	o.SetOnMessageHandler(func(client emitter.Emitter, msg emitter.Message) {
		fmt.Printf("Received message: %s\n", msg.Payload())
	})

	// Create a new emitter client and connect to the broker
	c := emitter.NewClient(o)
	sToken := c.Connect()
	if sToken.Wait() && sToken.Error() != nil {
		panic("Error on Client.Connect(): " + sToken.Error().Error())
	}

	// Subscribe to the presence demo channel
	c.Subscribe(key, channel)

	// Ask for presence
	r := emitter.NewPresenceRequest()
	r.Key = key
	r.Channel = channel
	c.Presence(r)

	var index int
	for range time.NewTicker(time.Second).C {
		// Publish to the channel
		index++
		c.Publish(key, channel, fmt.Sprintf("%d: example message", index))
	}
}
