package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/register"
	"ucenter/internal/svc"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *register.RegReq) (*register.RegRes, error) {
	// todo: add your logic here and delete this line

	return &register.RegRes{}, nil
}
