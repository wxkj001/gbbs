package user

import (
	"bbs/ent/user"
	"bbs/user/internal/svc"
	"bbs/user/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(GetUser *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	user := l.svcCtx.Ent.User.Query().Where(user.ID(1)).FirstX(l.ctx)
	return &types.GetUserResponse{
		NickName: user.NickName,
		Email:    user.Email,
		Phone:    user.Phone,
	}, nil
}
