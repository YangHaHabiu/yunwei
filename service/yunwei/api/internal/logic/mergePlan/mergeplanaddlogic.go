package mergePlan

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergePlanAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanAddLogic {
	return &MergePlanAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergePlanAddLogic) MergePlanAdd(req *types.AddMergePlanReq) error {
	var tmp []*yunweiclient.MergePlanCommon
	err := copier.Copy(&tmp, req.MergePlanData)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.MergePlanAdd(l.ctx, &yunweiclient.AddMergePlanReq{MergePlanData: tmp})
	if err != nil {
		return err
	}
	return nil
}
