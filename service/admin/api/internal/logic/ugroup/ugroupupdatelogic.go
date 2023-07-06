package ugroup

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUgroupUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupUpdateLogic {
	return &UgroupUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UgroupUpdateLogic) UgroupUpdate(req *types.UpdateUgroupReq) error {
	_, err := l.svcCtx.AdminRpc.UgroupUpdate(l.ctx, &admin.UgroupUpdateReq{
		Id:           req.Id,
		UgJson:       req.UgJson,
		UgName:       req.UgName,
		LastUpdateBy: ctxdata.GetUnameFromCtx(l.ctx),
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return nil
}
