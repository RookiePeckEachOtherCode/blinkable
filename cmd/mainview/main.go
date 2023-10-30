package main

import (
	mainview "blinkable/kitex_gen/Mainview/mainviewservice"
	"blinkable/pkg/viper"
	zlog "blinkable/pkg/zap"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap"
	"net"
)

var (
	cfg        = viper.Load("mainview_service").Viper
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
	MianviewServer := mainview.NewServer(
		new(MainviewServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serverName,
		}),
	)
	err = MianviewServer.Run()
	if err != nil {
		zap.S().Panicf("%v ===> %v", "server run", err)
	}
}
