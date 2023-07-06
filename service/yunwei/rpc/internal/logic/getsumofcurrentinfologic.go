package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSumOfCurrentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSumOfCurrentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSumOfCurrentInfoLogic {
	return &GetSumOfCurrentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Dashboard Start
func (l *GetSumOfCurrentInfoLogic) GetSumOfCurrentInfo(in *yunweiclient.GetSumOfCurrentInfoListReq) (*yunweiclient.GetSumOfCurrentInfoListResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "project_id__in", in.ProjectIds)

	list, err := l.svcCtx.StatServerGameInfoModel.FindPageListByPageGetBaseInfo(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询基础信息失败，原因：" + err.Error())
	}
	var tmp []*yunweiclient.GetSumOfCurrentInfoData
	err = copier.Copy(&tmp, list)
	if err != nil {
		return nil, xerr.NewErrMsg("复制基础信息失败，原因：" + err.Error())
	}
	return &yunweiclient.GetSumOfCurrentInfoListResp{
		Rows: tmp,
	}, nil
}
