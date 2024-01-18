package main

import (
	user "blinkable/server/kitex_gen/user/userservice"
	"blinkable/server/service/user/config"
	"blinkable/server/service/user/dao"
	i "blinkable/server/service/user/init"
	"context"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	i.InitLogger()
	r, info := i.InitNacos()
	db := i.InitDB()
	rdb := i.InitRedis()

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	impl := &UserServiceImpl{
		MysqlCen: dao.NewUser(db),
		RedisCen: dao.NewRedisCen(rdb),
	}

	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(config.GlobalServerConfig.Host, config.GlobalServerConfig.Port))
	if err != nil {
		klog.Errorf("resolve tcp addr :failed: %s", err)
	}
	svr := user.NewServer(
		impl,
		// server.WithServiceAddr(utils.NewNetAddr("tcp", net.JoinHostPort(config.GlobalServerConfig.Port, config.GlobalServerConfig.Port))),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
