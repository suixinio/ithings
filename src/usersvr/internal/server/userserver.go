// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/go-things/things/src/usersvr/internal/logic"
	"github.com/go-things/things/src/usersvr/internal/svc"
	"github.com/go-things/things/src/usersvr/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) RegisterCore(ctx context.Context, in *user.RegisterCoreReq) (*user.RegisterCoreResp, error) {
	l := logic.NewRegisterCoreLogic(ctx, s.svcCtx)
	return l.RegisterCore(in)
}

func (s *UserServer) Register2(ctx context.Context, in *user.Register2Req) (*user.Register2Resp, error) {
	l := logic.NewRegister2Logic(ctx, s.svcCtx)
	return l.Register2(in)
}

func (s *UserServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServer) GetUserCore(ctx context.Context, in *user.GetUserCoreReq) (*user.GetUserCoreResp, error) {
	l := logic.NewGetUserCoreLogic(ctx, s.svcCtx)
	return l.GetUserCore(in)
}

func (s *UserServer) CheckToken(ctx context.Context, in *user.CheckTokenReq) (*user.CheckTokenResp, error) {
	l := logic.NewCheckTokenLogic(ctx, s.svcCtx)
	return l.CheckToken(in)
}

func (s *UserServer) ModifyUserInfo(ctx context.Context, in *user.ModifyUserInfoReq) (*user.NilResp, error) {
	l := logic.NewModifyUserInfoLogic(ctx, s.svcCtx)
	return l.ModifyUserInfo(in)
}

func (s *UserServer) GetUserCoreList(ctx context.Context, in *user.GetUserCoreListReq) (*user.GetUserCoreListResp, error) {
	l := logic.NewGetUserCoreListLogic(ctx, s.svcCtx)
	return l.GetUserCoreList(in)
}
