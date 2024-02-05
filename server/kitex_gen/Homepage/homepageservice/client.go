// Code generated by Kitex v0.7.2. DO NOT EDIT.

package homepageservice

import (
	homepage "blinkable/server/kitex_gen/Homepage"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	LikeAction(ctx context.Context, req *homepage.LikeRequest, callOptions ...callopt.Option) (r *homepage.LikeResponse, err error)
	GetMainview(ctx context.Context, req *homepage.GetHomepageRequest, callOptions ...callopt.Option) (r *homepage.GetHomepageResponse, err error)
	AddGuestbook(ctx context.Context, req *homepage.AddGuestbookRequest, callOptions ...callopt.Option) (r *homepage.AddGuestbookResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kHomepageServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kHomepageServiceClient struct {
	*kClient
}

func (p *kHomepageServiceClient) LikeAction(ctx context.Context, req *homepage.LikeRequest, callOptions ...callopt.Option) (r *homepage.LikeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeAction(ctx, req)
}

func (p *kHomepageServiceClient) GetMainview(ctx context.Context, req *homepage.GetHomepageRequest, callOptions ...callopt.Option) (r *homepage.GetHomepageResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMainview(ctx, req)
}

func (p *kHomepageServiceClient) AddGuestbook(ctx context.Context, req *homepage.AddGuestbookRequest, callOptions ...callopt.Option) (r *homepage.AddGuestbookResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddGuestbook(ctx, req)
}
