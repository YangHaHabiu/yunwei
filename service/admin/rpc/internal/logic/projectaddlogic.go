package logic

import (
	"context"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectAddLogic {
	return &ProjectAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// project rpc start
func (l *ProjectAddLogic) ProjectAdd(in *adminclient.ProjectAddReq) (*adminclient.ProjectAddResp, error) {
	fmt.Println(gconv.Int64(in.ProjectType))
	obj, err := l.svcCtx.ProjectModel.Insert(l.ctx, &model.Project{
		ProjectCn:   in.ProjectCn,
		ProjectEn:   in.ProjectEn,
		ProjectTeam: in.ProjectTeam,
		ProjectType: gconv.Int64(in.ProjectType),
		GroupQq:     in.GroupQq,
		GroupType:   in.GroupType,
		GroupDevQq:  in.GroupDevQq,
	})

	if err != nil {
		return nil, xerr.NewErrMsg("插入项目数据失败，原因： " + err.Error())
	}
	id, err := obj.LastInsertId()
	if err != nil {
		return nil, xerr.NewErrMsg("查询插入项目记录失败，原因： " + err.Error())
	}
	err = l.svcCtx.ProjectRelationshipModel.Delete(l.ctx, id)
	if err != nil {
		return nil, xerr.NewErrMsg("删除项目与资产关联关系失败，原因： " + err.Error())
	}
	_, err = l.svcCtx.ProjectRelationshipModel.Insert(l.ctx, &model.ProjectRelationship{
		CompanyId: in.CompanyId,
		ProjectId: id,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("新增项目与资产关系失败，原因： " + err.Error())
	}

	return &adminclient.ProjectAddResp{}, nil
}
