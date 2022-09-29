//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

func initializeServer(ctx context.Context) server.IServer {
	wire.Build()

	return &server.Server{}
}
