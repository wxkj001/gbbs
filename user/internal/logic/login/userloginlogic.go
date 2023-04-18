package login

import (
	"bbs/common/errorx"
	"bbs/common/util"
	"bbs/ent"
	"bbs/ent/user"
	"bbs/user/internal/svc"
	"bbs/user/internal/types"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

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
	claims := make(jwt.MapClaims)
	claims["exp"] = l.svcCtx.Config.Auth.AccessExpire + time.Now().Unix()
	claims["iat"] = l.svcCtx.Config.Auth.AccessExpire
	claims["userUUId"] = first.ID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	to, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}
	//l.svcCtx.Ent.Topic.QueryUsers()
	ll := l.svcCtx.Ent.User.QueryTopics(first).WithUsers(func(query *ent.UserQuery) {
		query.Where(user.IDEQ(first.ID))
	}).AllX(l.ctx)
	l.Infof("%+v", ll)
	return &types.LoginResponse{Token: to}, nil
}
