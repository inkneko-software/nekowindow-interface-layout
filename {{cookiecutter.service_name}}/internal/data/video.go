package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type VideoRepo interface {
}

type videoRepo struct {
	data *Data
	log  *log.Helper
}

func NewIdentifyRepo(data *Data, logger log.Logger) VideoRepo {
	return &videoRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "package", "data")),
	}
}

func (r *videoRepo) ExampleRPCRequest(ctx context.Context, uid uint32, create int64, expire int64) (sessionKey string, err error) {
	/*
		resp, err:= r.data.videoSvc.SomeRpc(api.Req{})
		if err != nil{
			return "", err
		}
		return resp.data, nil
	*/
	return "", nil
}
