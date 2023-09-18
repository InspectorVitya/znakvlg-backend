package rest

import (
	"encoding/json"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/valyala/fasthttp"
	"net/http"
)

func (h *Handler) CreateUser(ctx *fasthttp.RequestCtx) {
	req := &model.Users{}
	if err := json.Unmarshal(ctx.Request.Body(), &req); err != nil {
		h.l.Errorf("Create user err:  %w", err)
		OutputJsonMessage(ctx, 500, ErrInternal.Error())
		return
	}
	req.Prepare()

	if invalid := h.userApp.ValidateUser(ctx, req); len(invalid) > 0 {
		OutputJson(ctx, http.StatusBadRequest, invalid)
		return
	}

	err := h.userApp.CreateUser(ctx, req)
	if err != nil {
		h.l.Errorf("Create user err:  %w", err)
		OutputJsonMessage(ctx, 500, err.Error())
		return
	}

	OutputJsonMessage(ctx, 201, "user created")
}

func (h *Handler) GetUserByID(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	user, err := h.userApp.GetUserByID(ctx, id)
	if err != nil {
		h.l.Errorf("GetUserByID err:  %w", err)
		OutputJsonMessage(ctx, 500, err.Error())
		return
	}

	OutputJson(ctx, 200, user)
}
