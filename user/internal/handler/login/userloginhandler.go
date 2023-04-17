package login

import (
	"bbs/common/response"
	"bbs/user/internal/logic/login"
	"bbs/user/internal/svc"
	"bbs/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(w, nil, err)
			return
		}

		l := login.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		response.Response(w, resp, err) //②
		/*if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}*/
	}
}
