package encoding

import "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/envelope"

// Marshaller are responsible to marshal Envelope object into specific formats (e.g a Sarama message)
type Marshaller interface {
	// Marshal a protobuffer message to specific format
	Marshal(t *envelope.Envelope, msg interface{}) error
}

// Unmarshaller are responsible to unmarshal input message into an envelope
type Unmarshaller interface {
	// Unmarshal high message into a Envelope
	Unmarshal(msg interface{}, t *envelope.Envelope) error
}
