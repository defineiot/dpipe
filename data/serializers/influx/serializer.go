package influx

import (
	"github.com/defineiot/dpipe/data/metrics"
)

type InfluxSerializer struct {
}

func (s *InfluxSerializer) Serialize(m metrics.Metric) ([]byte, error) {
	return m.Serialize(), nil
}
