package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTrendChartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTrendChartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTrendChartLogic {
	return &GetTrendChartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTrendChartLogic) GetTrendChart(in *yunweiclient.GetTrendChartListReq) (*yunweiclient.GetTrendChartListResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters,
		"project_id__in", in.ProjectIds,
		"count_type__=", in.Types,
	)
	list, err := l.svcCtx.StatServerGameInfoModel.FindPageListByPage(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询趋势信息失败，原因：" + err.Error())
	}
	var tmp []*yunweiclient.GetTrendChartData
	err = copier.Copy(&tmp, list)
	if err != nil {
		return nil, xerr.NewErrMsg("复制趋势信息失败，原因：" + err.Error())
	}
	return &yunweiclient.GetTrendChartListResp{
		Rows: tmp,
	}, nil
}
