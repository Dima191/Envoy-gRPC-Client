package app

import (
	"log/slog"
)

type serviceProvider struct {
	certServerPort  int
	routeServerPort int

	nodeID string

	protectedListenerPort   uint32
	unprotectedListenerPort uint32

	logger *slog.Logger
}

func newServiceProvider(certServerPort, routeServerPort int, nodeID string, protectedListenerPort, unprotectedListenerPort uint32, logger *slog.Logger) *serviceProvider {
	return &serviceProvider{
		certServerPort:          certServerPort,
		routeServerPort:         routeServerPort,
		protectedListenerPort:   protectedListenerPort,
		unprotectedListenerPort: unprotectedListenerPort,
		nodeID:                  nodeID,
		logger:                  logger,
	}
}
