package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptDeleteLogic) DeptDelete(in *adminclient.DeptDeleteReq) (*adminclient.DeptDeleteResp, error) {

	filters := make([]interface{}, 0)
	filters = append(filters, "dept_id__=", in.Id)
	//判断关联的用户
	all, err := l.svcCtx.UserModel.FindAll(l.ctx, &adminclient.UserListReq{}, filters...)
	if err != nil || len(*all) != 0 {
		return nil, xerr.NewErrMsg("部门关联用户数据，禁止删除，请检查")
	}
	//判断关联的项目
	filters = make([]interface{}, 0)
	filters = append(filters, "view_dept_id__=", in.Id)
	allx, err := l.svcCtx.ProjectModel.FindAll(l.ctx, filters...)
	if err != nil || len(*allx) != 0 {
		return nil, xerr.NewErrMsg("部门关联项目数据，禁止删除，请检查")
	}

	err = l.svcCtx.DeptModel.DeleteSoft(l.ctx, &model.SysDept{

		Id: in.Id,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_DELETE_ERROR)
	}

	return &adminclient.DeptDeleteResp{}, nil
}
