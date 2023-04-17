package user

import (
	"bbs/common/response"
	"bbs/user/internal/logic/user"
	"bbs/user/internal/svc"
	"bbs/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(w, nil, err)
			return
		}

		l := user.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		response.Response(w, resp, err) //②
		/*if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}*/
	}
}
