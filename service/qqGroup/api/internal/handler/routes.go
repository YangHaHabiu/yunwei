// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	report "ywadmin-v3/service/qqGroup/api/internal/handler/report"
	"ywadmin-v3/service/qqGroup/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/report",
				Handler: report.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/killProcess",
				Handler: report.KillProcessHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)
}
