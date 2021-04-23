package dsfr

import (
	"io"
	"net/url"
)

// ResourceFunc is a higher-level HTTP handler. Only GET requests without bodies are sent to resources.
type ResourceFunc func(io.Writer, url.Values) error

type Resources map[string]ResourceFunc
