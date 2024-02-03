package rpc

import (
	"blinkable/server/kitex_gen/Homepage/homepageservice"
	"blinkable/server/service/api/config"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	nacos "github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func initHomepage() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.GlobalNacosConfig.Host, config.GlobalNacosConfig.Port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         config.GlobalNacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}

	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}

	r := nacos.NewNacosResolver(nacosCli, nacos.WithGroup("HOMEPAGE_GROUP"))
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)

	c, err := homepageservice.NewClient(
		config.GlobalServerConfig.HomepageServer.Name,
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.HomepageServer.Name}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	config.GlobalHomepageClient = c
}
