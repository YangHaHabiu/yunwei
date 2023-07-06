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

type StgroupAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStgroupAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupAddLogic {
	return &StgroupAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StgroupAddLogic) StgroupAdd(req *types.AddStgroupReq) (err error) {
	_, err = l.svcCtx.AdminRpc.StgroupAdd(l.ctx, &admin.StgroupAddReq{
		StName:   req.StName,
		StJson:   req.StJson,
		StRemark: req.StRemark,
		CreateBy: ctxdata.GetUnameFromCtx(l.ctx),
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	//刷新策略
	common.FlushStrategy(l.svcCtx, l.ctx, ctxdata.GetUidFromCtx(l.ctx), ctxdata.GetUnameFromCtx(l.ctx))

	return nil

}
