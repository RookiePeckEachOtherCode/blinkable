// Code generated by hertz generator.

package main

import (
	"blinkable/server/service/api/config"
	i "blinkable/server/service/api/init"
	"blinkable/server/service/api/init/rpc"
	"context"
	"fmt"

	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/pprof"
	"github.com/kitex-contrib/obs-opentelemetry/provider"

	"github.com/cloudwego/hertz/pkg/app/server"
	hertzTracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
)

func main() {
	i.InitLogger()
	r, info := i.InitNacos()
	i.InitMinio()
	// i.InitCreateALLBucket()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	tracer, cfg := hertzTracing.NewServerTracer()
	rpc.Init()

	h := server.New(
		tracer,
		server.WithHostPorts(fmt.Sprintf(":%s", config.GlobalServerConfig.Port)),
		server.WithRegistry(r, info),
		server.WithHandleMethodNotAllowed(true),
	)

	pprof.Register(h)
	h.Use(hertzTracing.ServerMiddleware(cfg))
	h.Use(gzip.Gzip(gzip.DefaultCompression))
	register(h)
	h.Spin()
}
