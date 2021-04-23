package example

import (
	"context"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func (s *ExampleDataSource) StatsQueryHandler(ctx context.Context, query backend.DataQuery) (data.Framer, error) {
	return nil, nil
}

func (s *ExampleDataSource) TableQueryHandler(ctx context.Context, query backend.DataQuery) (data.Framer, error) {
	return nil, nil
}
