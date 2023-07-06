package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCheckStategyInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCheckStategyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCheckStategyInfoLogic {
	return &GetUserCheckStategyInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCheckStategyInfoLogic) GetUserCheckStategyInfo(in *adminclient.StgroupUserCheckInfoReq) (*adminclient.StgroupUserCheckInfoResp, error) {

	allugroup, err := l.svcCtx.StgroupUgroupModel.FindAll(l.ctx, "stgroup_id__=", in.Id)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询策略关联用户组失败")
	}
	alluser, err := l.svcCtx.StgroupUserModel.FindAll(l.ctx, "stgroup_id__=", in.Id)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询策略关联用户失败")
	}
	userCheck := make([]*adminclient.DataCheck, 0)
	ugroupCheck := make([]*adminclient.DataCheck, 0)
	for _, v := range *alluser {
		userCheck = append(userCheck, &adminclient.DataCheck{Id: v.UserId})
	}
	for _, v := range *allugroup {
		ugroupCheck = append(ugroupCheck, &adminclient.DataCheck{Id: v.UgroupId})
	}

	return &adminclient.StgroupUserCheckInfoResp{
		UserCheck:   userCheck,
		UgroupCheck: ugroupCheck,
	}, nil
}
