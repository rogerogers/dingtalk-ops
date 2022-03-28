package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/rogerogers/dingtalk-ops/api/helloworld/v1"
	"github.com/rogerogers/dingtalk-ops/internal/biz"
)

type dingtalkRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewDingtalkRepo(data *Data, logger log.Logger) biz.DingtalkRepo {
	return &dingtalkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *dingtalkRepo) GetUserToken(ctx context.Context, d *v1.GetUserTokenRequest) (*v1.GetUserTokenReply, error) {
	return &v1.GetUserTokenReply{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}, nil
}

func (r *dingtalkRepo) GetUserInfoByToken(ctx context.Context, d *v1.GetUserInfoByTokenRequest) (*v1.GetUserInfoByTokenReply, error) {
	return &v1.GetUserInfoByTokenReply{}, nil
}
func (r *dingtalkRepo) GetUserIdByUnionId(ctx context.Context, d *v1.GetUserIdByUnionIdRequest) (*v1.GetUserIdByUnionIdReply, error) {
	return &v1.GetUserIdByUnionIdReply{}, nil
}

func (r *dingtalkRepo) GetUserInfoByUserId(ctx context.Context, d *v1.GetUserInfoByUserIdRequest) (*v1.GetUserInfoByUserIdReply, error) {
	return &v1.GetUserInfoByUserIdReply{}, nil
}
