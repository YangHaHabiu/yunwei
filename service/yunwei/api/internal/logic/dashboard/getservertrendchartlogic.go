package dashboard

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerTrendChartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServerTrendChartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerTrendChartLogic {
	return &GetServerTrendChartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServerTrendChartLogic) GetServerTrendChart(req *types.GetServerTrendChartListReq) (resp *types.GetServerTrendChartListResp, err error) {
	//个人项目值及列表
	projectIds, _, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	list, err := l.svcCtx.YunWeiRpc.GetTrendChart(l.ctx, &yunweiclient.GetTrendChartListReq{ProjectIds: projectIds, Types: "server"})
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.GetServerTrendChartData, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制信息出错，原因：" + err.Error())
	}
	resp = new(types.GetServerTrendChartListResp)
	resp.Rows = tmp
	return
}
