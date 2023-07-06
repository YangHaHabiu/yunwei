package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMaintainPlanGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanGetLogic {
	return &MaintainPlanGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MaintainPlanGetLogic) MaintainPlanGet(in *yunweiclient.GetMaintainPlanReq) (*yunweiclient.ListMaintainPlanData, error) {
	//one, err := l.svcCtx.MaintainPlanModel.FindOne(l.ctx, in.Id)
	//if err != nil {
	//	return nil, xerr.NewErrMsg("查询单条数据失败")
	//}
	all, err := l.svcCtx.MaintainPlanModel.FindAll(l.ctx, "id__=", in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}

	if len(*all) != 1 {
		return nil, xerr.NewErrMsg("存在多条或者不存在此数据")
	}

	var tmp yunweiclient.ListMaintainPlanData
	err = copier.Copy(&tmp, (*all)[0])
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
