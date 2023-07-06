package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectUpdateLogic {
	return &ProjectUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProjectUpdateLogic) ProjectUpdate(in *adminclient.ProjectUpdateReq) (*adminclient.ProjectUpdateResp, error) {
	err := l.svcCtx.ProjectModel.Update(l.ctx, &model.Project{
		ProjectId:   in.ProjectId,
		ProjectCn:   in.ProjectCn,
		ProjectEn:   in.ProjectEn,
		ProjectTeam: in.ProjectTeam,
		ProjectType: gconv.Int64(in.ProjectType),
		GroupQq:     in.GroupQq,
		GroupType:   in.GroupType,
		GroupDevQq:  in.GroupDevQq,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	one, err := l.svcCtx.ProjectRelationshipModel.FindOneByPrId(l.ctx, in.ProjectId)
	if err != nil {
		_, err = l.svcCtx.ProjectRelationshipModel.Insert(l.ctx, &model.ProjectRelationship{
			CompanyId: in.CompanyId,
			ProjectId: in.ProjectId,
		})
		if err != nil {
			return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
		}
	} else {
		if one.CompanyId != in.CompanyId {
			err = l.svcCtx.ProjectRelationshipModel.Update(l.ctx, &model.ProjectRelationship{
				Id:        one.Id,
				CompanyId: in.CompanyId,
				ProjectId: in.ProjectId,
			})
			if err != nil {
				return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
			}
		}
	}

	return &adminclient.ProjectUpdateResp{}, nil
}
