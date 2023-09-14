package rest

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type out struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PagingList struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func OutputJson(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	// Marshal provided interface into JSON structure
	jsonResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		OutputJsonMessage(ctx, 500, "json parsing error")
		return
	}
	// Write content-type, statuscode, payload
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.SetStatusCode(code)
	fmt.Fprint(ctx, string(jsonResult))
	ctx.Response.Header.Set("Connection", "close")
}

func OutputJsonMessage(ctx *fasthttp.RequestCtx, code int, msg string) {
	// Write content-type, statuscode, payload
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.SetStatusCode(code)
	o := out{code, msg}
	jsonResult, _ := json.Marshal(o)
	fmt.Fprint(ctx, string(jsonResult))
	ctx.Response.Header.Set("Connection", "close")
}

func OutputPagingList(ctx *fasthttp.RequestCtx, code int, total int64, list interface{}) {
	// Write content-type, statuscode, payload
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "_t")
	ctx.Response.Header.SetStatusCode(code)
	out := PagingList{total, list}
	jsonResult, _ := json.Marshal(out)
	fmt.Fprint(ctx, string(jsonResult))
	ctx.Response.Header.Set("Connection", "close")
}

func OutputCORSOptions(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "text/html")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Access-Control-Allow-Methods,Content-Type")
	ctx.Response.Header.SetStatusCode(200)
	ctx.Response.Header.Set("Connection", "close")
}
