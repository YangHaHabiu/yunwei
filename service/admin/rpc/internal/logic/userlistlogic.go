package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *adminclient.UserListReq) (*adminclient.UserListResp, error) {
	var (
		count int64
		list  []*adminclient.UserListData
		err   error
		all   *[]model.SysUserList
	)

	filters := make([]interface{}, 0)
	filters = append(filters, "su.id__=",
		in.UserId,
		"su.name__regexp",
		in.Name,
		"su.nick_name__regexp",
		in.NickName,
		"su.email__like",
		in.Email,
		"su.status__=",
		in.Status,
		"su.mobile__like",
		in.Mobile,
		"su.dept_id__in",
		in.DeptIds)
	count, _ = l.svcCtx.UserModel.Count(l.ctx, in, filters...)
	if in.Current != 0 && in.PageSize != 0 {
		all, err = l.svcCtx.UserModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, in, filters...)

	} else {
		all, err = l.svcCtx.UserModel.FindAll(l.ctx, in, filters...)

	}
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询用户列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询用户列表信息失败")
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("复制用户列表信息失败")
	}

	return &adminclient.UserListResp{
		Total: count,
		List:  list,
	}, nil
}
