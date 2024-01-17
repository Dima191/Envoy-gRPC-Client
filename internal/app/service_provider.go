package app

import (
	"context"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"log/slog"
	xdscache "xds_server/internal/cache"
	certclient "xds_server/internal/client/cert"
	certclientimpl "xds_server/internal/client/cert/implementation"
	routesclient "xds_server/internal/client/routes"
	routesclientimpl "xds_server/internal/client/routes/implementation"
	xdsservice "xds_server/internal/service"
)

type serviceProvider struct {
	cdsCache *cache.LinearCache
	ldsCache *cache.LinearCache
	rdsCache *cache.LinearCache

	srv xdsservice.Service

	certCl       certclient.Client
	routesClient routesclient.Client

	certServerPort  int
	routeServerPort int

	nodeID string

	protectedListenerPort   uint32
	unprotectedListenerPort uint32

	logger *slog.Logger
}

func (sp *serviceProvider) CertClient() (certclient.Client, error) {
	if sp.certCl == nil {
		var err error
		sp.certCl, err = certclientimpl.New(sp.certServerPort, sp.logger)
		if err != nil {
			return nil, err
		}
	}
	return sp.certCl, nil
}

func (sp *serviceProvider) RoutesClient() (routesclient.Client, error) {
	if sp.routesClient == nil {
		var err error
		sp.routesClient, err = routesclientimpl.New(sp.routeServerPort, sp.logger)
		if err != nil {
			return nil, err
		}
	}
	return sp.routesClient, nil
}

func (sp *serviceProvider) Cache(ctx context.Context) (*cache.LinearCache, *cache.LinearCache, *cache.LinearCache, error) {
	if sp.cdsCache == nil || sp.ldsCache == nil || sp.rdsCache == nil {
		certClient, err := sp.CertClient()
		if err != nil {
			return nil, nil, nil, err
		}

		routesClient, err := sp.RoutesClient()
		if err != nil {
			return nil, nil, nil, err
		}

		sp.cdsCache, sp.ldsCache, sp.rdsCache, err = xdscache.InitializeCache(ctx, certClient, routesClient, sp.protectedListenerPort, sp.unprotectedListenerPort, sp.logger)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	return sp.cdsCache, sp.ldsCache, sp.rdsCache, nil
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
