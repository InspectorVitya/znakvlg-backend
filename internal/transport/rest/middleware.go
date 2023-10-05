package rest

import (
	"github.com/valyala/fasthttp"
	"strings"
)

func (h *Handler) AuthAdmin(next func(ctx *fasthttp.RequestCtx, userID string)) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Cookie("accessToken"))
		if strings.Compare(token, "") == 0 {
			h.l.Error("Authorization cookie is required")
			ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
			return
		}

		payload, err := h.authApp.AuthAdmin(ctx, token)
		if err != nil {
			h.l.Error(err)
			ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
			return
		}
		next(ctx, payload.UserID)
		return

	})
}
