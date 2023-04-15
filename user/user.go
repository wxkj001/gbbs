package user

import (
	"bbs/config"
	"bbs/ent"
	"bbs/user/internal/handler"
	"bbs/user/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func NewUserServie(c config.Config, server *rest.Server, db *ent.Client) {
	ctx := svc.NewServiceContext(c, db)
	handler.RegisterHandlers(server, ctx)
}
