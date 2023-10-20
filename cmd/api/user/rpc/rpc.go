package rpc

import (
	"blinkable/common/errno"
	"blinkable/kitex_gen/user"
	"blinkable/kitex_gen/user/userservice"
	"blinkable/pkg/viper"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	cfg        = viper.Load("user_service")
	serverName = cfg.Viper.GetString("server.name")
	etcdAddr   = fmt.Sprintf("%s:%d", cfg.Viper.GetString("etcd.host"), cfg.Viper.GetInt("etcd.port"))
)

var userClient userservice.Client

func init() {
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	errno.HandleErrWithPanic("用户客户端初始化失败", err)

	_userClient, err := userservice.NewClient(
		serverName,
		client.WithRPCTimeout(30*time.Second),        // 设置RPC超时时间
		client.WithConnectTimeout(36000*time.Second), // 设置连接超时时间
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serverName,
		}),
	)

	if err != nil {
		errno.HandleErrWithPanic("用户客户端初始化失败", err)
	}
	userClient = _userClient
}

func Login(ctx context.Context, c *app.RequestContext) {
}

func Register(ctx context.Context, req *user.UserRegisterRequest) (*user.UseeRegisterResponse, error) {
	return userClient.UserRegister(ctx, req)
}
