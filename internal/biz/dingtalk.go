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

func (uc *DingtalkUseCase) GetUserInfoByToken(ctx context.Context, d *v1.GetUserInfoByTokenRequest) (*v1.GetUserInfoByTokenReply, error) {
	uc.log.WithContext(ctx).Infof("getUserInfoByToken: %v", d.AccessToken)
	return uc.repo.GetUserInfoByToken(ctx, d)
}

func (uc *DingtalkUseCase) GetUserIdByUnionId(ctx context.Context, d *v1.GetUserIdByUnionIdRequest) (*v1.GetUserIdByUnionIdReply, error) {
	uc.log.WithContext(ctx).Infof("getUserIdByUnionId: %v", d.UnionId)
	return uc.repo.GetUserIdByUnionId(ctx, d)
}

func (uc *DingtalkUseCase) GetUserInfoByUserId(ctx context.Context, d *v1.GetUserInfoByUserIdRequest) (*v1.GetUserInfoByUserIdReply, error) {
	uc.log.WithContext(ctx).Infof("getUserInfoByUserId: %v", d.UserId)
	return uc.repo.GetUserInfoByUserId(ctx, d)
}
