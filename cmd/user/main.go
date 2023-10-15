package main

import (
	user "blinkable/cmd/user/kitex_gen/user/userservice"
	"blinkable/common/consts"
	"blinkable/errs"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCD_ADDR})

	if err != nil {
		log.Panicf("用户服务注册失败: %s", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", consts.UserServiceAPIPort)

	errs.HandleErrWithPanic("用户服务API端口解析失败", err)

	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: consts.UserServiceName,
		}),
	)

	err = svr.Run()

	errs.HandleErrWithNormal("用户服务启动失败", err)
}
