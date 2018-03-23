package processors

import (
	"github.com/defineiot/dpipe/data/metrics"
)

// Processor use to handler data
type Processor interface {
	// SampleConfig returns the default configuration of the Input
	SampleConfig() string

	// Description returns a one-sentence description on the Input
	Description() string

	// Apply the filter to the given metric
	Apply(in ...metrics.Metric) []metrics.Metric
}
