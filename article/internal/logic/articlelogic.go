package logic

import (
	"context"

	"bbs/article/internal/svc"
	"bbs/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleLogic) Article(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
