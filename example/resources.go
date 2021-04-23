package example

import (
	"encoding/json"
	"io"
	"net/url"
)

type Label struct {
	Name  string
	Color string
}

// ListLabels is a resourcefunc that calls the example API and returns a list of labels
func (s *ExampleDataSource) ListLabels(w io.Writer, v url.Values) error {
	return json.NewEncoder(w).Encode(Label{
		Name: "example",
	})
}
