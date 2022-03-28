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
