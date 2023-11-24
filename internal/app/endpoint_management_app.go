package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"
)

func (a *App) initEMDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{}

	for _, f := range deps {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runEMGRPCServer() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.EndpointManagementPort))
	if err != nil {
		a.logger.Error("failed to create connection for endpoint_management app", slog.String("error", err.Error()))
		return err
	}
	a.logger.Info("endpoint_management app started", slog.Uint64("port", uint64(a.cfg.EndpointManagementPort)))

	if err = a.emGRPC.Serve(listener); err != nil {
		a.logger.Error("failed to serve endpoint_management app", slog.String("error", err.Error()))
		return err
	}
	return nil
}
