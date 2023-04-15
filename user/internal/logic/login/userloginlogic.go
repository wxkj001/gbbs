package login

import (
	"bbs/common/errorx"
	"bbs/common/util"
	"bbs/ent/user"
	"bbs/user/internal/svc"
	"bbs/user/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	l.Debugf("login:%+v", req)
	first, err := l.svcCtx.Ent.User.Query().Where(user.Or(user.EmailEQ(req.Name), user.PhoneEQ(req.Name))).First(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(4004, "user not found")
	}
	l.Debugf("user:%+v", first)
	u := util.Util{Config: l.svcCtx.Config}
	if !u.PassVerify(req.Password, first.Password) {
		return nil, errorx.NewCodeError(4005, "password not")
	}
	return &types.LoginResponse{Token: first.UserUUID.String()}, nil
}
