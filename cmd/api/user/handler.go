package userHandler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	// "blinkable/cmd/user/kitex_gen/user/userservice"
	// "blinkable/common/consts"
	// "blinkable/errs"
	// "context"
	// "github.com/cloudwego/hertz/pkg/app"
	// "github.com/cloudwego/kitex/client"
	// etcd "github.com/kitex-contrib/registry-etcd"
)

// var userClinet userservice.Client

func init() {
	// r, err := etcd.NewEtcdResolver([]string{consts.ETCD_ADDR})
	// errs.HandleErrWithPanic("用户客户端初始化失败", err)

	// userClinet, err := userservice.NewClient(
	// 	consts.UserServiceName,
	// 	client.WithResolver(r),
	// )

	// errs.HandleErrWithPanic("用户客户端初始化失败", err)
}

func Login(ctx context.Context, c *app.RequestContext) {
}
func Register(ctx context.Context, c *app.RequestContext) {
}
