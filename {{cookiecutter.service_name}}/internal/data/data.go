package data

import (
	"context"
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/conf"
	videov1 "nekowindow-backend/app/service/main/video/api/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewVideoClient)

// Data .
type Data struct {
	// TODO wrapped database client
	videoSvc videov1.VideoClient
}

// NewData .
func NewData(
	c *conf.Data,
	videoSvc videov1.VideoClient,
	logger log.Logger,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{videoSvc: videoSvc}, cleanup, nil
}

func NewVideoClient(r registry.Discovery) videov1.VideoClient {

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery://default/nekowindow.service.main.video"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := videov1.NewVideoClient(conn)
	return c
}
