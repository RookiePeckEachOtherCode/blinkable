// Code generated by Kitex v0.7.3. DO NOT EDIT.

package mainviewservice

import (
	mainview "blinkable/kitex_gen/Mainview"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return mainviewServiceServiceInfo
}

var mainviewServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MainviewService"
	handlerType := (*mainview.MainviewService)(nil)
	methods := map[string]kitex.MethodInfo{
		"LikeAction":   kitex.NewMethodInfo(likeActionHandler, newMainviewServiceLikeActionArgs, newMainviewServiceLikeActionResult, false),
		"GetMainview":  kitex.NewMethodInfo(getMainviewHandler, newMainviewServiceGetMainviewArgs, newMainviewServiceGetMainviewResult, false),
		"AddGuestbook": kitex.NewMethodInfo(addGuestbookHandler, newMainviewServiceAddGuestbookArgs, newMainviewServiceAddGuestbookResult, false),
		"ChangeCard":   kitex.NewMethodInfo(changeCardHandler, newMainviewServiceChangeCardArgs, newMainviewServiceChangeCardResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "mainview",
		"ServiceFilePath": `idl/Mainview.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func likeActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*mainview.MainviewServiceLikeActionArgs)
	realResult := result.(*mainview.MainviewServiceLikeActionResult)
	success, err := handler.(mainview.MainviewService).LikeAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMainviewServiceLikeActionArgs() interface{} {
	return mainview.NewMainviewServiceLikeActionArgs()
}

func newMainviewServiceLikeActionResult() interface{} {
	return mainview.NewMainviewServiceLikeActionResult()
}

func getMainviewHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*mainview.MainviewServiceGetMainviewArgs)
	realResult := result.(*mainview.MainviewServiceGetMainviewResult)
	success, err := handler.(mainview.MainviewService).GetMainview(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMainviewServiceGetMainviewArgs() interface{} {
	return mainview.NewMainviewServiceGetMainviewArgs()
}

func newMainviewServiceGetMainviewResult() interface{} {
	return mainview.NewMainviewServiceGetMainviewResult()
}

func addGuestbookHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*mainview.MainviewServiceAddGuestbookArgs)
	realResult := result.(*mainview.MainviewServiceAddGuestbookResult)
	success, err := handler.(mainview.MainviewService).AddGuestbook(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMainviewServiceAddGuestbookArgs() interface{} {
	return mainview.NewMainviewServiceAddGuestbookArgs()
}

func newMainviewServiceAddGuestbookResult() interface{} {
	return mainview.NewMainviewServiceAddGuestbookResult()
}

func changeCardHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*mainview.MainviewServiceChangeCardArgs)
	realResult := result.(*mainview.MainviewServiceChangeCardResult)
	success, err := handler.(mainview.MainviewService).ChangeCard(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMainviewServiceChangeCardArgs() interface{} {
	return mainview.NewMainviewServiceChangeCardArgs()
}

func newMainviewServiceChangeCardResult() interface{} {
	return mainview.NewMainviewServiceChangeCardResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) LikeAction(ctx context.Context, req *mainview.LikeRequest) (r *mainview.LikeResponse, err error) {
	var _args mainview.MainviewServiceLikeActionArgs
	_args.Req = req
	var _result mainview.MainviewServiceLikeActionResult
	if err = p.c.Call(ctx, "LikeAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMainview(ctx context.Context, req *mainview.GetMainvewRequest) (r *mainview.GetMainviewResponse, err error) {
	var _args mainview.MainviewServiceGetMainviewArgs
	_args.Req = req
	var _result mainview.MainviewServiceGetMainviewResult
	if err = p.c.Call(ctx, "GetMainview", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddGuestbook(ctx context.Context, req *mainview.AddGuestbookRequest) (r *mainview.AddGuestbookResponse, err error) {
	var _args mainview.MainviewServiceAddGuestbookArgs
	_args.Req = req
	var _result mainview.MainviewServiceAddGuestbookResult
	if err = p.c.Call(ctx, "AddGuestbook", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeCard(ctx context.Context, req *mainview.ChangeCardRequest) (r *mainview.ChangeCardResponse, err error) {
	var _args mainview.MainviewServiceChangeCardArgs
	_args.Req = req
	var _result mainview.MainviewServiceChangeCardResult
	if err = p.c.Call(ctx, "ChangeCard", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
