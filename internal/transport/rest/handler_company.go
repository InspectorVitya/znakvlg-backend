package rest

import (
	"encoding/json"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
)

func (h *Handler) CreateCompany(ctx *fasthttp.RequestCtx) {
	req := model.CompanyRequest{}
	contentHeader := ctx.Request.Header.Peek("Content-Type")
	//var filesHeaders []*multipart.FileHeader
	switch {
	case strings.Contains(string(contentHeader), "multipart/form-data"):
		form, err := ctx.MultipartForm()
		if err != nil {
			h.l.Error("failed CreateCompany get MultipartForm err: ", err.Error())
			OutputJsonMessage(ctx, 400, err.Error())
			return
		}
		err = json.Unmarshal([]byte(form.Value["data"][0]), &req)
		if err != nil {
			h.l.Error("failed CreateCompany marshal json err: ", err.Error())
			OutputJsonMessage(ctx, 400, err.Error())
			return
		}
		//todo  доделать загрузку файлов
		//filesHeaders = form.File["avatar"]

	case strings.Contains(string(contentHeader), "application/json"):
		err := json.Unmarshal(ctx.Request.Body(), &req)
		if err != nil {
			h.l.Error("failed CreateCompany marshal json err: ", err.Error())
			OutputJsonMessage(ctx, 400, err.Error())
			return
		}
	}
	err := h.companyApp.CreateCompany(ctx, req.Company, req.PlatesUse)
	if err != nil {
		h.l.Error("failed CreateCompany: ", err.Error())
		OutputJsonMessage(ctx, 400, err.Error())
		return
	}
	OutputJsonMessage(ctx, 201, "company created")
}

func (h *Handler) GetCompany(ctx *fasthttp.RequestCtx) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		h.l.Error("failed GetCompany: ", err.Error())
		OutputJsonMessage(ctx, 400, err.Error())
		return
	}
	company, err := h.companyApp.GetCompanyByID(ctx, uint32(id))
	if err != nil {
		h.l.Error("failed GetCompany: ", err.Error())
		OutputJsonMessage(ctx, 400, err.Error())
		return
	}
	OutputJson(ctx, 200, company)
}
