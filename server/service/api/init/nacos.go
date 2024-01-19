package init

import (
	"blinkable/server/common/consts"
	"blinkable/server/service/api/config"
	"net"

	"github.com/bwmarrin/snowflake"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

// InitNacos to init nacos
func InitNacos() (registry.Registry, *registry.Info) {
	v := viper.New()
	v.SetConfigFile(consts.APIConfigPath)
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed: %s", err)
	}
	if err := v.Unmarshal(&config.GlobalNacosConfig); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err)
	}
	klog.Infof("Gloabl nacos's config Info : %s", config.GlobalNacosConfig)

	// Read configuration information from nacos
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
		Username:            config.GlobalServerConfig.Name,
		Password:            config.GlobalNacosConfig.Password,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		klog.Fatalf("create config client failed: %s", err.Error())
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: config.GlobalNacosConfig.DataId,
		Group:  config.GlobalNacosConfig.Group,
	})
	if err != nil {
		klog.Fatalf("get config info failed: %s", err.Error())
	}

	err = sonic.Unmarshal([]byte(content), &config.GlobalServerConfig)
	if err != nil {
		klog.Fatalf("nacos config failed: %s", err.Error())
	}

	registryClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		klog.Fatal("create registryClient err: %s", err.Error())
	}

	r := nacos.NewNacosRegistry(registryClient, nacos.WithRegistryGroup(config.GlobalNacosConfig.Group))

	sf, err := snowflake.NewNode(consts.NacosSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate nacos service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: config.GlobalServerConfig.Name,
		Addr:        utils.NewNetAddr("tcp", net.JoinHostPort(config.GlobalServerConfig.Host, config.GlobalServerConfig.Port)),
		Tags:        map[string]string{"ID": sf.Generate().Base36()},
		Weight:      registry.DefaultWeight,
	}

	return r, info
}
