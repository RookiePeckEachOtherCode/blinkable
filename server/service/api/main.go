// Code generated by hertz generator.

package main

import (
	"blinkable/server/service/api/config"
	i "blinkable/server/service/api/init"
	"blinkable/server/service/api/init/rpc"
	"context"
	"fmt"
	"time"

	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/pprof"
	"github.com/kitex-contrib/obs-opentelemetry/provider"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
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
	h.Use(cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))
	register(h)
	h.Spin()
}
