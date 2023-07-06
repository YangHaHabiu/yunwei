package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserStrategyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserStrategyListLogic {
	return &UserStrategyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserStrategyListLogic) UserStrategyList(in *adminclient.UserStrategyInfoReq) (*adminclient.UserStrategyInfoResp, error) {
	userObj, err := l.svcCtx.UserModel.SelectStrategyInfoByUname(l.ctx, in.Name)
	if err != nil || len(*userObj) != 1 {
		return nil, xerr.NewErrMsg("查询用户相关的策略信息失败")
	}

	return &adminclient.UserStrategyInfoResp{
		StgroupStJson: (*userObj)[0].StgroupStJson,
		SysUserId:     (*userObj)[0].SysUserId,
		SysUserName:   (*userObj)[0].SysUserName,
	}, nil
}
