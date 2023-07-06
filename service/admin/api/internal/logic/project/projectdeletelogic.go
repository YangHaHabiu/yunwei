package project

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectDeleteLogic {
	return &ProjectDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectDeleteLogic) ProjectDelete(req *types.DeleteProjectReq) (err error) {
	var delFlag int64
	one, err := l.svcCtx.AdminRpc.ProjectGetOne(l.ctx, &admin.ProjectGetOneReq{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return err
	}
	if one.DelFlag == 0 {
		delFlag = 1
	}

	_, err = l.svcCtx.AdminRpc.ProjectDelete(l.ctx, &admin.ProjectDeleteReq{
		ProjectId: req.ProjectId,
		DelFlag:   delFlag,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除异常:%s", req.ProjectId, err.Error())
		return err
	}
	return
}
