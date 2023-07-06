package mergePlan

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergePlanUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanUpdateLogic {
	return &MergePlanUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergePlanUpdateLogic) MergePlanUpdate(req *types.UpdateMergePlanReq) error {
	var tmp yunwei.MergePlanCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝数据失败，原因：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.MergePlanUpdate(l.ctx, &yunweiclient.UpdateMergePlanReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
