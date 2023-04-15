package handler

import (
	"net/http"

	"bbs/article/internal/logic"
	"bbs/article/internal/svc"
	"bbs/article/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), svcCtx)
		resp, err := l.Article(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
