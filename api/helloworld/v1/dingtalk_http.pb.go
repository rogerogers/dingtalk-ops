// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type DingtalkHTTPServer interface {
	GetUserIdByUnionId(context.Context, *GetUserIdByUnionIdRequest) (*GetUserIdByUnionIdReply, error)
	GetUserInfoByToken(context.Context, *GetUserInfoByTokenRequest) (*GetUserInfoByTokenReply, error)
	GetUserInfoByUserId(context.Context, *GetUserInfoByUserIdRequest) (*GetUserInfoByUserIdReply, error)
	GetUserToken(context.Context, *GetUserTokenRequest) (*GetUserTokenReply, error)
}

func RegisterDingtalkHTTPServer(s *http.Server, srv DingtalkHTTPServer) {
	r := s.Route("/")
	r.GET("/dingtalk/token-by-auth-code", _Dingtalk_GetUserToken0_HTTP_Handler(srv))
	r.GET("/dingtalk/user-info-by-token", _Dingtalk_GetUserInfoByToken0_HTTP_Handler(srv))
	r.GET("/dingtalk/unionid-by-userid", _Dingtalk_GetUserIdByUnionId0_HTTP_Handler(srv))
	r.GET("/dingtalk/user-info-by-userid", _Dingtalk_GetUserInfoByUserId0_HTTP_Handler(srv))
}

func _Dingtalk_GetUserToken0_HTTP_Handler(srv DingtalkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.Dingtalk/GetUserToken")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserToken(ctx, req.(*GetUserTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserTokenReply)
		return ctx.Result(200, reply)
	}
}

func _Dingtalk_GetUserInfoByToken0_HTTP_Handler(srv DingtalkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoByTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.Dingtalk/GetUserInfoByToken")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfoByToken(ctx, req.(*GetUserInfoByTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoByTokenReply)
		return ctx.Result(200, reply)
	}
}

func _Dingtalk_GetUserIdByUnionId0_HTTP_Handler(srv DingtalkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserIdByUnionIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.Dingtalk/GetUserIdByUnionId")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserIdByUnionId(ctx, req.(*GetUserIdByUnionIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserIdByUnionIdReply)
		return ctx.Result(200, reply)
	}
}

func _Dingtalk_GetUserInfoByUserId0_HTTP_Handler(srv DingtalkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoByUserIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.Dingtalk/GetUserInfoByUserId")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfoByUserId(ctx, req.(*GetUserInfoByUserIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoByUserIdReply)
		return ctx.Result(200, reply)
	}
}

type DingtalkHTTPClient interface {
	GetUserIdByUnionId(ctx context.Context, req *GetUserIdByUnionIdRequest, opts ...http.CallOption) (rsp *GetUserIdByUnionIdReply, err error)
	GetUserInfoByToken(ctx context.Context, req *GetUserInfoByTokenRequest, opts ...http.CallOption) (rsp *GetUserInfoByTokenReply, err error)
	GetUserInfoByUserId(ctx context.Context, req *GetUserInfoByUserIdRequest, opts ...http.CallOption) (rsp *GetUserInfoByUserIdReply, err error)
	GetUserToken(ctx context.Context, req *GetUserTokenRequest, opts ...http.CallOption) (rsp *GetUserTokenReply, err error)
}

type DingtalkHTTPClientImpl struct {
	cc *http.Client
}

func NewDingtalkHTTPClient(client *http.Client) DingtalkHTTPClient {
	return &DingtalkHTTPClientImpl{client}
}

func (c *DingtalkHTTPClientImpl) GetUserIdByUnionId(ctx context.Context, in *GetUserIdByUnionIdRequest, opts ...http.CallOption) (*GetUserIdByUnionIdReply, error) {
	var out GetUserIdByUnionIdReply
	pattern := "/dingtalk/unionid-by-userid"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.Dingtalk/GetUserIdByUnionId"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DingtalkHTTPClientImpl) GetUserInfoByToken(ctx context.Context, in *GetUserInfoByTokenRequest, opts ...http.CallOption) (*GetUserInfoByTokenReply, error) {
	var out GetUserInfoByTokenReply
	pattern := "/dingtalk/user-info-by-token"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.Dingtalk/GetUserInfoByToken"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DingtalkHTTPClientImpl) GetUserInfoByUserId(ctx context.Context, in *GetUserInfoByUserIdRequest, opts ...http.CallOption) (*GetUserInfoByUserIdReply, error) {
	var out GetUserInfoByUserIdReply
	pattern := "/dingtalk/user-info-by-userid"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.Dingtalk/GetUserInfoByUserId"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DingtalkHTTPClientImpl) GetUserToken(ctx context.Context, in *GetUserTokenRequest, opts ...http.CallOption) (*GetUserTokenReply, error) {
	var out GetUserTokenReply
	pattern := "/dingtalk/token-by-auth-code"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.Dingtalk/GetUserToken"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
