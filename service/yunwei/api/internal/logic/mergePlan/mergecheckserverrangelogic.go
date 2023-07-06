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

type MergeCheckServerRangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergeCheckServerRangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergeCheckServerRangeLogic {
	return &MergeCheckServerRangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergeCheckServerRangeLogic) MergeCheckServerRange(req *types.MergeCheckServerRangeReq) (resp *types.MergeCheckServerRangeResp, err error) {
	var tmp yunwei.MergePlanCommon
	err = copier.Copy(&tmp, req)
	if err != nil {
		return nil, xerr.NewErrMsg("检查合服范围拷贝数据失败，原因：" + err.Error())
	}

	info, err := l.svcCtx.YunWeiRpc.MergeCheckServerRange(l.ctx, &yunweiclient.MergeCheckServerRangeReq{One: &tmp})
	if err != nil {
		return nil, err
	}
	resp = new(types.MergeCheckServerRangeResp)
	resp.CombineRange = info.CombineRange
	return
}
