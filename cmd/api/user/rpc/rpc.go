package rpc

import (
	"blinkable/kitex_gen/user"
	"blinkable/kitex_gen/user/userservice"
	"blinkable/pkg/viper"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap"
)

var (
	cfg        = viper.Load("user_service")
	serverName = cfg.Viper.GetString("server.name")
	etcdAddr   = fmt.Sprintf("%s:%d", cfg.Viper.GetString("etcd.host"), cfg.Viper.GetInt("etcd.port"))
)

var userClient userservice.Client

func init() {
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		zap.S().Panicf("%v ===> %v", "etcd resolver", err)
	}

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
		zap.S().Panicf("%v ===> %v", "用户客户端初始化失败", err)
	}
	userClient = _userClient
}

func Login(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	return userClient.UserLogin(ctx, req)
}

func Register(ctx context.Context, req *user.UserRegisterRequest) (*user.UseeRegisterResponse, error) {
	return userClient.UserRegister(ctx, req)
}

func Info(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	return userClient.UserInfo(ctx, req)
}

func UserInfoUpdate(ctx context.Context, req *user.UserInfoUpdateRequest) (*user.UserInfoUpdateResponse, error) {
	return userClient.UserInfoUpdate(ctx, req)
}
