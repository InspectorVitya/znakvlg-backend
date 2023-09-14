package rest

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/app"
	"github.com/InspectorVitya/znakvlg-backend/pkg/logger"
	"github.com/fasthttp/router"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var (
	ErrInternal   = errors.New("Server is down")
	ErrBadGateway = errors.New("Bad gateway")
	ErrNotFound   = errors.New("Not found")
	ErrAuth       = errors.New("Authorization header required")
	ErrTokenParse = errors.New("Token parse error")
)

type Handler struct {
	serverHost string
	server     *fasthttp.Server
	service    *app.App
	l          *logger.Logger
}

func New(s *app.App, serverHost string, log *logger.Logger) *Handler {
	h := new(Handler)
	h.serverHost = serverHost
	h.service = s
	h.l = log
	return h
}

func (h *Handler) Run() error {
	h.l.Info("Start app")

	r := router.New()

	r.GET("/", HealthCheck)

	admin := r.Group("/admin")
	h.initAdminRouter(admin)

	h.l.Info("Run app HTTP server on adr: ", h.serverHost)
	server := &fasthttp.Server{
		Handler:            r.Handler,
		MaxRequestBodySize: 1 * 1024 * 1024,
	}

	return server.ListenAndServe(h.serverHost)
}

func (h *Handler) initAdminRouter(r *router.Group) {
	company := r.Group("/company")
	{
		company.GET("/{id}", h.GetCompany)
		company.POST("/", h.CreateCompany)
	}
	user := r.Group("/user")
	{
		user.POST("/", h.CreateUser)
		user.GET("/{id}", h.GetUserByID)
	}
}

// HealthCheck - get the status of server.
func HealthCheck(ctx *fasthttp.RequestCtx) {
	OutputJson(ctx, 200, map[string]interface{}{
		"data": "Server is up and running",
	})
}

func (h *Handler) Stop(ctx context.Context) error {
	return h.server.ShutdownWithContext(ctx)
}
