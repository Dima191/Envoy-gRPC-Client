package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"
)

func (a *App) initXDSDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{}

	for _, f := range deps {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runXDSGRPCServer() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.XdsPort))
	if err != nil {
		a.logger.Error("failed to create connection for xds app", slog.String("error", err.Error()))
		return err
	}

	a.logger.Info("xds app started", slog.Uint64("port", uint64(a.cfg.XdsPort)))
	if err = a.xdsGRPC.Serve(listener); err != nil {
		a.logger.Error("failed to serve xds app", slog.String("error", err.Error()))
		return err
	}

	return nil
}
