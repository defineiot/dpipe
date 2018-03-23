package aggregators

type Creator func() Aggregator

var Aggregators = map[string]Creator{}

func Add(name string, creator Creator) {
	Aggregators[name] = creator
}
