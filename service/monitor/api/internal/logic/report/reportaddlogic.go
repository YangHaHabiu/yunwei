package report

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/monitor/rpc/monitorclient"

	"ywadmin-v3/service/monitor/api/internal/svc"
	"ywadmin-v3/service/monitor/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReportAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportAddLogic {
	return &ReportAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReportAddLogic) ReportAdd(req *types.ReportAddReq) error {
	var tmp monitorclient.ReportAddReq
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("复制参数失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.Monitor.ReportAdd(l.ctx, &tmp)
	if err != nil {
		return err
	}

	return nil
}
