package serviceimpl

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"sync"
	certclient "xds_server/internal/client/cert"
	xdsservice "xds_server/internal/service"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

type service struct {
	cdsCache *cache.LinearCache
	ldsCache *cache.LinearCache
	rdsCache *cache.LinearCache

	mu sync.Mutex

	v *validator.Validate

	certClient certclient.Client

	nodeID string

	protectedListenerPort   uint32
	unprotectedListenerPort uint32

	logger *slog.Logger
}

func New(cdsCache, ldsCache, rdsCache *cache.LinearCache, certClient certclient.Client, nodeID string, protectedListenerPort, unprotectedListenerPort uint32, logger *slog.Logger) xdsservice.Service {
	return &service{
		cdsCache:                cdsCache,
		ldsCache:                ldsCache,
		rdsCache:                rdsCache,
		v:                       validator.New(),
		certClient:              certClient,
		nodeID:                  nodeID,
		protectedListenerPort:   protectedListenerPort,
		unprotectedListenerPort: unprotectedListenerPort,
		logger:                  logger,
	}
}
