package rpc

import (
	mainview "blinkable/kitex_gen/Mainview"
	"blinkable/kitex_gen/Mainview/mainviewservice"
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
	cfg        = viper.Load("mainview_service")
	serverName = cfg.Viper.GetString("server.name")
	etcdAddr   = fmt.Sprintf("%s:%d", cfg.Viper.GetString("etcd.host"), cfg.Viper.GetInt("etcd.port"))
)
var mainviewClient mainviewservice.Client

func init() {
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})

	if err != nil {
		zap.S().Panicf("%v ===> %v", "etcd resolver", err)
	}

	_mainviewClient, err := mainviewservice.NewClient(
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
	mainviewClient = _mainviewClient
}

func Getmainview(ctx context.Context, req *mainview.GetMainvewRequest) (*mainview.GetMainviewResponse, error) {
	return mainviewClient.GetMainview(ctx, req)
}

func LikeAction(ctx context.Context, req *mainview.LikeRequest) (*mainview.LikeResponse, error) {
	return mainviewClient.LikeAction(ctx, req)
}

func AddGuestbook(ctx context.Context, req *mainview.AddGuestbookRequest) (*mainview.AddGuestbookResponse, error) {
	return mainviewClient.AddGuestbook(ctx, req)
}
func ChangeCard(ctx context.Context, req *mainview.ChangeCardRequest) (*mainview.ChangeCardResponse, error) {
	return mainviewClient.ChangeCard(ctx, req)
}
