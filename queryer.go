package dsfr

import (
	"context"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type QueryFunc func(context.Context, backend.DataQuery) (data.Framer, error)

type Queries map[string]QueryFunc
