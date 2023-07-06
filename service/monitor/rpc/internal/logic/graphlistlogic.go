package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/monitor/rpc/internal/svc"
	"ywadmin-v3/service/monitor/rpc/monitorclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GraphListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGraphListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GraphListLogic {
	return &GraphListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GraphListLogic) GraphList(in *monitorclient.GraphListReq) (*monitorclient.GraphListResp, error) {
	byType, err := l.svcCtx.ReportStreamMinuteModel.SelectListAllByType(l.ctx, in)
	if err != nil {
		return nil, xerr.NewErrMsg("查询上报信息失败，原因：" + err.Error())
	}
	marshal, err := json.Marshal(byType)
	if err != nil {
		return nil, xerr.NewErrMsg("生成信息失败，原因：" + err.Error())

	}
	return &monitorclient.GraphListResp{
		Rows: string(marshal),
	}, nil
}
