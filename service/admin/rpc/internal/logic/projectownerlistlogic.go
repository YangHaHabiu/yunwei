package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/gogf/gf/util/gconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectOwnerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectOwnerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectOwnerListLogic {
	return &ProjectOwnerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProjectOwnerListLogic) ProjectOwnerList(in *adminclient.ProjectOwnerReq) (*adminclient.ProjectListResp, error) {
	allOwner, err := l.svcCtx.ProjectModel.FindListByUserId(l.ctx, in.UserId, in.ProjectType)
	if err != nil {
		return nil, xerr.NewErrMsg("根据用户ID查询个人项目信息失败")
	}
	var list []*adminclient.ProjectListData

	for _, data := range *allOwner {
		list = append(list, &adminclient.ProjectListData{
			ProjectId:   data.ProjectId,
			ProjectCn:   data.ProjectCn,
			ProjectEn:   data.ProjectEn,
			ProjectTeam: data.ProjectTeam,
			ProjectType: gconv.String(data.ProjectType),
			GroupQq:     data.GroupQq,
			GroupType:   data.GroupType,
			GroupDevQq:  data.GroupDevQq,
		})
	}
	return &adminclient.ProjectListResp{
		List: list,
	}, nil
}
