package publisher

import (
	"context"
	"errors"
)

const PluginName = "publisher"

type PublisherServer struct {
	Impl Publisher
	UnimplementedPublisherPluginServer
}

func (p *PublisherServer) Init(ctx context.Context, request *PublisherInit_Request) (*PublisherInit_Response, error) {
	err := p.Impl.Init(request.Config)
	res := &PublisherInit_Response{}
	if err != nil {
		res.Error = err.Error()
	}
	return res, nil
}

func (p *PublisherServer) Name(ctx context.Context, request *PublisherName_Request) (*PublisherName_Response, error) {
	return &PublisherName_Response{Name: p.Impl.Name()}, nil
}

func (p *PublisherServer) Version(ctx context.Context, request *PublisherVersion_Request) (*PublisherVersion_Response, error) {
	return &PublisherVersion_Response{Version: p.Impl.Version()}, nil
}

func (p *PublisherServer) Publish(ctx context.Context, request *PublisherPublish_Request) (*PublisherPublish_Response, error) {
	err := p.Impl.Publish(request.NewRelease)
	res := &PublisherPublish_Response{}
	if err != nil {
		res.Error = err.Error()
	}
	return res, nil
}

type PublisherClient struct {
	Impl PublisherPluginClient
}

func (c *PublisherClient) Init(m map[string]string) error {
	res, err := c.Impl.Init(context.Background(), &PublisherInit_Request{
		Config: m,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (c *PublisherClient) Name() string {
	res, err := c.Impl.Name(context.Background(), &PublisherName_Request{})
	if err != nil {
		panic(err)
	}
	return res.Name
}

func (c *PublisherClient) Version() string {
	res, err := c.Impl.Version(context.Background(), &PublisherVersion_Request{})
	if err != nil {
		panic(err)
	}
	return res.Version
}

func (c *PublisherClient) Publish(newRelease string) error {
	res, err := c.Impl.Publish(context.Background(), &PublisherPublish_Request{
		NewRelease: newRelease,
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}
