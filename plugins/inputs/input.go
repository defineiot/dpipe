package inputs

import (
	"github.com/defineiot/dpipe/data"
)

// Input is collect data to input
type Input interface {
	// SampleConfig returns the default configuration of the Input
	SampleConfig() string

	// Description returns a one-sentence description on the Input
	Description() string

	// Gather takes in an accumulator and adds the metrics that the Input
	// gathers. This is called every "interval"
	Gather(data.Accumulator) error
}

// ServiceInput is via a service to input data
type ServiceInput interface {
	// SampleConfig returns the default configuration of the Input
	SampleConfig() string

	// Description returns a one-sentence description on the Input
	Description() string

	// Gather takes in an accumulator and adds the metrics that the Input
	// gathers. This is called every "interval"
	Gather(data.Accumulator) error

	// Start starts the ServiceInput's service, whatever that may be
	Start(data.Accumulator) error

	// Stop stops the services and closes any necessary channels and connections
	Stop()
}
