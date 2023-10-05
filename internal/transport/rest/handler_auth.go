package rest

import (
	"encoding/json"
	"errors"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/valyala/fasthttp"
	"net/http"
)

func (h *Handler) SignIn(ctx *fasthttp.RequestCtx) {
	authUser := model.RequestAuth{}
	if err := json.Unmarshal(ctx.Request.Body(), &authUser); err != nil {
		h.l.Error("failed SignIn marshal json err: ", err.Error())
		OutputJsonMessage(ctx, 400, err.Error())
		return
	}
	token, refresh, err := h.authApp.SignIn(ctx, authUser)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUserBlocked):
			OutputJsonMessage(ctx, http.StatusForbidden, "Пользователь заблокирован.")
		case errors.Is(err, model.ErrInvalidPassword):
			OutputJsonMessage(ctx, http.StatusForbidden, "Проверьте правильность написания логина и пароля.")
		default:
			OutputJsonMessage(ctx, http.StatusForbidden, "Произошла ошибка на сервере.")
		}

		return
	}
	h.l.Infof("domain: %s", ctx.Request.Host)
	if token != "" && refresh != "" {
		ctx.Response.Header.SetCookie(createCookie("refreshToken", refresh, 5184000))
		ctx.Response.Header.SetCookie(createCookie("accessToken", token, 5184000))
	} else {
		OutputJsonMessage(ctx, 500, "no token or refresh token")
		return
	}

	OutputJsonMessage(ctx, 200, "auth")
}

func (h *Handler) Logout(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.DelCookie("refreshToken")
	ctx.Response.Header.DelCookie("accessToken")

	OutputJsonMessage(ctx, 200, "logout")
}
