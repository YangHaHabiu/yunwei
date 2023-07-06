package maintainPlan

import (
	"context"
	"time"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMaintanListByPriLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMaintanListByPriLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMaintanListByPriLogic {
	return &GetMaintanListByPriLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMaintanListByPriLogic) GetMaintanListByPri(req *types.GetMaintanListByPriReq) (resp *types.GetMaintanListByPriResp, err error) {

	//获取昨日凌晨时间戳
	ts := time.Now().AddDate(0, 0, -1)
	timeYesterDay := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location()).Unix()
	list, err := l.svcCtx.YunWeiRpc.MaintainPlanList(l.ctx, &yunweiclient.ListMaintainPlanReq{
		ProjectId: req.ProjectId,
		StartTime: timeYesterDay,
		Current:   0,
		PageSize:  0,
		//MaintainType: "1,3,4",
		TaskId: -1,
	})
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListMaintainPlanData, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制维护计划失败，原因：" + err.Error())
	}

	resp = new(types.GetMaintanListByPriResp)
	resp.Rows = tmp
	return
}
