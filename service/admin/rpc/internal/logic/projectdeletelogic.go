package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectDeleteLogic {
	return &ProjectDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProjectDeleteLogic) ProjectDelete(in *adminclient.ProjectDeleteReq) (*adminclient.ProjectDeleteResp, error) {
	err := l.svcCtx.ProjectModel.DeleteSoft(l.ctx, in.ProjectId, in.DelFlag)
	if err != nil {
		return nil, xerr.NewErrMsg("操作项目失败，原因：" + err.Error())
	}

	return &adminclient.ProjectDeleteResp{}, nil
}
