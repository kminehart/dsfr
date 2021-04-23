package dsfr

import (
	"errors"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
)

type DataSource interface {
	Queries() Queries
	Resources() Resources
	HealthCheckers() []HealthChecker
	Reload(backend.DataSourceInstanceSettings) (DataSource, error)
}

type Handler struct {
	im instancemgmt.InstanceManager
}

func (h *Handler) CallResource(ctx backend.PluginContext, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) CheckHealth(ctx backend.PluginContext, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	instance, err := h.im.Get(ctx)
	if err != nil {
		return nil, err
	}
	ds, ok := instance.(DataSource)

	if !ok {
		return nil, errors.New("invalid datasource")
	}

	checks := ds.HealthCheckers()

	return nil, nil
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
//
// The Frames' RefID property, when it is an empty string, will be automatically set to
// the RefID in QueryDataResponse.Responses map. This is done before the QueryDataResponse is
// sent to Grafana. Therefore one does not need to be set that property on frames when using this method.
func (h *Handler) QueryData(ctx backend.PluginContext, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	panic("not implemented") // TODO: Implement
}
