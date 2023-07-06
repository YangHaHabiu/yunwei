package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogAddLogic {
	return &SysLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// log rpc start
func (l *SysLogAddLogic) SysLogAdd(in *adminclient.SysLogAddReq) (*adminclient.SysLogAddResp, error) {
	_, err := l.svcCtx.SysLogModel.Insert(l.ctx, &model.SysLog{
		UserName:  in.UserName,
		Operation: in.Operation,
		Method:    in.Method,
		Params:    in.Params,
		Time:      in.Time,
		Ip:        in.Ip,
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}

	return &adminclient.SysLogAddResp{}, nil
}
