// Code generated by goctl. DO NOT EDIT!
// Source: dm.proto

package server

import (
	"context"

	"gitee.com/godLei6/things/src/dmsvr/dm"
	"gitee.com/godLei6/things/src/dmsvr/internal/logic"
	"gitee.com/godLei6/things/src/dmsvr/internal/svc"
)

type DmServer struct {
	svcCtx *svc.ServiceContext
}

func NewDmServer(svcCtx *svc.ServiceContext) *DmServer {
	return &DmServer{
		svcCtx: svcCtx,
	}
}

func (s *DmServer) LoginAuth(ctx context.Context, in *dm.LoginAuthReq) (*dm.Response, error) {
	l := logic.NewLoginAuthLogic(ctx, s.svcCtx)
	return l.LoginAuth(in)
}

func (s *DmServer) AccessAuth(ctx context.Context, in *dm.AccessAuthReq) (*dm.Response, error) {
	l := logic.NewAccessAuthLogic(ctx, s.svcCtx)
	return l.AccessAuth(in)
}

func (s *DmServer) ManageDevice(ctx context.Context, in *dm.ManageDeviceReq) (*dm.DeviceInfo, error) {
	l := logic.NewManageDeviceLogic(ctx, s.svcCtx)
	return l.ManageDevice(in)
}

func (s *DmServer) ManageProduct(ctx context.Context, in *dm.ManageProductReq) (*dm.ProductInfo, error) {
	l := logic.NewManageProductLogic(ctx, s.svcCtx)
	return l.ManageProduct(in)
}

func (s *DmServer) GetProductInfo(ctx context.Context, in *dm.GetProductInfoReq) (*dm.ProductInfo, error) {
	l := logic.NewGetProductInfoLogic(ctx, s.svcCtx)
	return l.GetProductInfo(in)
}

func (s *DmServer) GetDeviceInfo(ctx context.Context, in *dm.GetDeviceInfoReq) (*dm.DeviceInfo, error) {
	l := logic.NewGetDeviceInfoLogic(ctx, s.svcCtx)
	return l.GetDeviceInfo(in)
}
