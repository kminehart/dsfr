package dsfr

import "context"

// HealthChecker checks the health of a service
type HealthChecker interface {
	CheckHealth(context.Context) error
}
