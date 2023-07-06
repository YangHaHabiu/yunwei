package inws

import (
	"context"
	"net/http"

	"ywadmin-v3/service/inws/api/internal/svc"
	"ywadmin-v3/service/inws/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInsideTasksLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewGetInsideTasksLogLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *GetInsideTasksLogLogic {
	return &GetInsideTasksLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *GetInsideTasksLogLogic) GetInsideTasksLog(req *types.GetInsideTasksLogReq) {

	ServeWs(l.w, l.r, l.svcCtx, req)
}
