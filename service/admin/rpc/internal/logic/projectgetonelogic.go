package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectGetOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectGetOneLogic {
	return &ProjectGetOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProjectGetOneLogic) ProjectGetOne(in *adminclient.ProjectGetOneReq) (*adminclient.ProjectListData, error) {

	one, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.ProjectId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条项目数据失败，原因：" + err.Error())
	}
	var tmp adminclient.ProjectListData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("复制单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
