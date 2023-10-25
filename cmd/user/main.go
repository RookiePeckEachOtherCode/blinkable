package main

import (
	user "blinkable/kitex_gen/user/userservice"
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap"
)

var (
	cfg        = viper.Load("user_service").Viper
	serverName = cfg.GetString("server.name")
	serverAddr = fmt.Sprintf("%s:%d", cfg.GetString("server.host"), cfg.GetInt("server.port"))
	etcdAddr   = fmt.Sprintf("%s:%d", cfg.GetString("etcd.host"), cfg.GetInt("etcd.port"))
)

func main() {
	zlog.Init()

	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})

	if err != nil {
		zap.S().Panicf("%v ===> %v", "etcd registry", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", serverAddr)

	if err != nil {
		zap.S().Panicf("%v ===> %v", "server host port", err)
	}

	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serverName,
		}),
	)

	err = svr.Run()

	if err != nil {
		zap.S().Panicf("%v ===> %v", "server run", err)
	}
}
