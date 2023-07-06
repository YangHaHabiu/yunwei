package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *adminclient.RoleListReq) (*adminclient.RoleListResp, error) {
	var (
		all *[]model.SysRole
		err error
	)
	count, _ := l.svcCtx.RoleModel.Count(l.ctx)

	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.RoleModel.FindAll(l.ctx)
	} else {
		all, err = l.svcCtx.RoleModel.FindPageListByPage(l.ctx, in.Current, in.PageSize)
	}
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询角色列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrCode(xerr.ADMIN_ROLESELECT_ERROR)
	}

	var list []*adminclient.RoleListData
	for _, role := range *all {
		list = append(list, &adminclient.RoleListData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        role.DelFlag,
		})
	}

	return &adminclient.RoleListResp{
		Total: count,
		List:  list,
	}, nil

}
