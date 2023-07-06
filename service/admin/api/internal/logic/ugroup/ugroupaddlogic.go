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

type UgroupAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUgroupAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupAddLogic {
	return &UgroupAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UgroupAddLogic) UgroupAdd(req *types.AddUgroupReq) (err error) {
	_, err = l.svcCtx.AdminRpc.UgroupAdd(l.ctx, &admin.UgroupAddReq{
		UgName:   req.UgName,
		UgJson:   req.UgJson,
		CreateBy: ctxdata.GetUnameFromCtx(l.ctx),
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return nil
}
