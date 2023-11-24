package certclient

import (
	"context"
	xdsmodels "xds_server/internal/models"
)

type Client interface {
	Cert(ctx context.Context, domain string) (*xdsmodels.Cert, error)
	CloseConn() error
}
