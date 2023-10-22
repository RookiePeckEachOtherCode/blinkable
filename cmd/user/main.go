package main

import (
	"blinkable/common/errno"
	user "blinkable/kitex_gen/user/userservice"
	"blinkable/pkg/viper"
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	cfg        = viper.Load("user_service")
	serverName = cfg.Viper.GetString("server.name")
	serverAddr = fmt.Sprintf("%s:%d", cfg.Viper.GetString("server.host"), cfg.Viper.GetInt("server.port"))
	etcdAddr   = fmt.Sprintf("%s:%d", cfg.Viper.GetString("etcd.host"), cfg.Viper.GetInt("etcd.port"))
)

func main() {
	where := "server_user"
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})

	if err != nil {
		errno.HandleErrWithPanic(where, "etcd registry", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", serverAddr)

	if err != nil {
		errno.HandleErrWithPanic(where, "server host port", err)
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
		errno.HandleErrWithPanic(where, "server run", err)
	}
}
