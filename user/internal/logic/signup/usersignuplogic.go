package signup

import (
	"bbs/common/errorx"
	"bbs/common/util"
	"context"
	"github.com/google/uuid"

	"bbs/user/internal/svc"
	"bbs/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignupLogic {
	return &UserSignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignupLogic) UserSignup(req *types.SignupRequest) error {
	u := util.Util{Config: l.svcCtx.Config}
	us, err := l.svcCtx.Ent.User.Create().
		SetGroupUUID(uuid.MustParse(l.svcCtx.Config.User.OrdinaryUUID)).
		SetNickName(req.NickName).
		SetEmail(req.Email).
		SetPhone(req.Phone).
		SetPassword(u.NewPassword(req.Password)).Save(l.ctx)
	if err != nil {
		return errorx.NewCodeError(4006, err.Error())
	}
	l.Infof("User:%+v", us)
	return nil
}
