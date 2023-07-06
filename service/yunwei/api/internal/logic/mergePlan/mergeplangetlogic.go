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

type MergePlanGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergePlanGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanGetLogic {
	return &MergePlanGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergePlanGetLogic) MergePlanGet(req *types.GetMergePlanReq) (*types.ListMergePlanData, error) {
	get, err := l.svcCtx.YunWeiRpc.MergePlanGet(l.ctx, &yunweiclient.GetMergePlanReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var tmp types.ListMergePlanData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
