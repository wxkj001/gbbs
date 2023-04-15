package signup

import (
	"bbs/common/response"
	"bbs/user/internal/logic/signup"
	"bbs/user/internal/svc"
	"bbs/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserSignupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignupRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(w, nil, err)
			return
		}

		l := signup.NewUserSignupLogic(r.Context(), svcCtx)
		err := l.UserSignup(&req)
		response.Response(w, nil, err) //②
		/*if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}*/
	}
}
