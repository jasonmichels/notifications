package main

import (    
    "github.com/jasonmichels/notifications/transports"
)

func main() {
    transportResolver := new(transports.TransportResolver)

    // Get a new transport instance that sends notification
    // This needs upated to add the data from the queue
    transport := transportResolver.ResolveTransport()
    transport.Send()
}