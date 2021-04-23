package example

import (
	"github.com/grafana/dsfr"
)

// QueryTypes are used to describe different types of queries
const (
	QueryTypeTable = "table"
	QueryTypeStats = "stats"
)

const (
	ResourcePathLabels = "/labels"
)

type ExampleDataSource struct{}

func (e *ExampleDataSource) Queries() dsfr.Queries {
	return dsfr.Queries{
		QueryTypeTable: e.TableQueryHandler,
		QueryTypeStats: e.TableQueryHandler,
	}
}

func (e *ExampleDataSource) Resources() dsfr.Resources {
	return dsfr.Resources{
		ResourcePathLabels: e.ListLabels,
	}
}

func (e *ExampleDataSource) HealthCheckers() []dsfr.HealthChecker {
	panic("not implemented") // TODO: Implement
}
