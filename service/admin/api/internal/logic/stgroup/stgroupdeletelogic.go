package stgroup

import (
	"context"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/api/internal/logic/common"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type StgroupDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStgroupDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupDeleteLogic {
	return &StgroupDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StgroupDeleteLogic) StgroupDelete(req *types.DeleteStgroupReq) (err error) {
	_, err = l.svcCtx.AdminRpc.StgroupDelete(l.ctx, &admin.StgroupDeleteReq{
		Id: req.StgroupId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除异常:%s", req.StgroupId, err.Error())
		return err
	}
	//刷新策略
	common.FlushStrategy(l.svcCtx, l.ctx, ctxdata.GetUidFromCtx(l.ctx), ctxdata.GetUnameFromCtx(l.ctx))
	return nil
}
