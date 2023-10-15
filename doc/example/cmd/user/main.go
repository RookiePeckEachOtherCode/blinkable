package main

import (
	user "example/cmd/user/kitex_gen/user/userservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	//使用etcd作为注册中心
	//etcd的地址为: 127.0.0.1:2379
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})

	if err != nil {
		panic(err)
	}

	//使用8888端口作为user服务端口
	addr, err := net.ResolveTCPAddr("tcp", ":8888")

	if err != nil {
		panic(err)
	}

	svr := user.NewServer(
		new(UserServiceImpl),         //new UserService接口的实现
		server.WithRegistry(r),       //使用etcd(变量r)作为注册中心
		server.WithServiceAddr(addr), //使用8888端口作为user服务端口
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ //设置服务的基本信息
			ServiceName: "user-service", //服务名
		}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
