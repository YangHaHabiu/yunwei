package stgroup

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/api/internal/logic/common"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type StgroupUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStgroupUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupUpdateLogic {
	return &StgroupUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StgroupUpdateLogic) StgroupUpdate(req *types.UpdateStgroupReq) error {
	_, err := l.svcCtx.AdminRpc.StgroupUpdate(l.ctx, &admin.StgroupUpdateReq{
		Id:           req.Id,
		StName:       req.StName,
		StRemark:     req.StRemark,
		StJson:       req.StJson,
		LastUpdateBy: ctxdata.GetUnameFromCtx(l.ctx),
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	//刷新策略
	common.FlushStrategy(l.svcCtx, l.ctx, ctxdata.GetUidFromCtx(l.ctx), ctxdata.GetUnameFromCtx(l.ctx))

	return nil
}
