package outputs

import (
	"github.com/defineiot/dpipe/data/metrics"
)

// Output put the data
type Output interface {
	// Connect to the Output
	Connect() error
	// Close any connections to the Output
	Close() error
	// Description returns a one-sentence description on the Output
	Description() string
	// SampleConfig returns the default configuration of the Output
	SampleConfig() string
	// Write takes in group of points to be written to the Output
	Write(metrics []metrics.Metric) error
}

// ServiceOutput put the data to an service
type ServiceOutput interface {
	// Connect to the Output
	Connect() error
	// Close any connections to the Output
	Close() error
	// Description returns a one-sentence description on the Output
	Description() string
	// SampleConfig returns the default configuration of the Output
	SampleConfig() string
	// Write takes in group of points to be written to the Output
	Write(metrics []metrics.Metric) error
	// Start the "service" that will provide an Output
	Start() error
	// Stop the "service" that will provide an Output
	Stop()
}
