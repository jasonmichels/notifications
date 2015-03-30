package transports

import (   
)

type Transporter interface {
    Send()
}

type TransportResolver struct {

}

// Resolve the correct message transporter from the data that will get passed in
func (resolver *TransportResolver) ResolveTransport() Transporter {
	// Actually do logic based on the data to pick correct transporter, if others exist
	return new(EmailTransport)
}