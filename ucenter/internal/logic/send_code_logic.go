package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/register"
	"ucenter/internal/svc"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *register.CodeReq) (*register.NoRes, error) {
	// todo: add your logic here and delete this line

	return &register.NoRes{}, nil
}
