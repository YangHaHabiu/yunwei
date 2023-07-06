package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLabelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelAddLogic {
	return &LabelAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// label rpc start
func (l *LabelAddLogic) LabelAdd(in *adminclient.LabelAddReq) (*adminclient.LabelAddResp, error) {
	_, err := l.svcCtx.LabelModel.Insert(l.ctx, &model.Label{
		LabelName:   in.LabelName,
		LabelValues: in.LabelValues,
		LabelRemark: in.LabelRemark,
		LabelType:   in.LabelType,
		CreateBy:    in.CreateBy,
		DelFlag:     0,
	})

	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	return &adminclient.LabelAddResp{}, nil
}
