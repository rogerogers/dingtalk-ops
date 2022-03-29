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
	return &v1.GetUserTokenReply{AccessToken: in.AuthCode, RefreshToken: g.RefreshToken}, nil
}

func (s *DingtalkService) GetUserInfoByUserId(ctx context.Context, in *v1.GetUserInfoByUserIdRequest) (*v1.GetUserInfoByUserIdReply, error) {
	g, err := s.uc.GetUserInfoByUserId(ctx, in)
	if err != nil {
		return nil, err
	}
	return g, nil
}
