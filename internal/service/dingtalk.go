package service

import (
	"context"

	v1 "github.com/rogerogers/dingtalk-ops/api/helloworld/v1"
	"github.com/rogerogers/dingtalk-ops/internal/biz"
)

type DingtalkService struct {
	v1.UnimplementedDingtalkServer

	uc *biz.DingtalkUseCase
}

func NewDingtalkService(uc *biz.DingtalkUseCase) *DingtalkService {
	return &DingtalkService{uc: uc}
}

func (s *DingtalkService) GetUserToken(ctx context.Context, in *v1.GetUserTokenRequest) (*v1.GetUserTokenReply, error) {
	g, err := s.uc.GetUserToken(ctx, in)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s *DingtalkService) GetUserInfoByToken(ctx context.Context, in *v1.GetUserInfoByTokenRequest) (*v1.GetUserInfoByTokenReply, error) {
	g, err := s.uc.GetUserInfoByToken(ctx, in)
	if err != nil {
		return nil, err
	}
	return g, nil
}
func (s *DingtalkService) GetUserIdByUnionId(ctx context.Context, in *v1.GetUserIdByUnionIdRequest) (*v1.GetUserIdByUnionIdReply, error) {
	g, err := s.uc.GetUserIdByUnionId(ctx, in)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s *DingtalkService) GetUserInfoByUserId(ctx context.Context, in *v1.GetUserInfoByUserIdRequest) (*v1.GetUserInfoByUserIdReply, error) {
	g, err := s.uc.GetUserInfoByUserId(ctx, in)
	if err != nil {
		return nil, err
	}
	return g, nil
}
