package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/rogerogers/dingtalk-ops/api/helloworld/v1"
	"golang.org/x/net/context"
)

type DingtalkRepo interface {
	GetUserToken(context.Context, *v1.GetUserTokenRequest) (*v1.GetUserTokenReply, error)
	GetUserInfoByToken(context.Context, *v1.GetUserInfoByTokenRequest) (*v1.GetUserInfoByTokenReply, error)
	GetUserIdByUnionId(context.Context, *v1.GetUserIdByUnionIdRequest) (*v1.GetUserIdByUnionIdReply, error)
	GetUserInfoByUserId(context.Context, *v1.GetUserInfoByUserIdRequest) (*v1.GetUserInfoByUserIdReply, error)
}

type DingtalkUseCase struct {
	repo DingtalkRepo
	log  *log.Helper
}

func NewDingtalkUseCase(repo DingtalkRepo, logger log.Logger) *DingtalkUseCase {
	return &DingtalkUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DingtalkUseCase) GetUserToken(ctx context.Context, d *v1.GetUserTokenRequest) (*v1.GetUserTokenReply, error) {
	uc.log.WithContext(ctx).Infof("getUserToken: %v", d.AuthCode)
	return uc.repo.GetUserToken(ctx, d)
}
