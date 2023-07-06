package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/model"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideOperationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideOperationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideOperationAddLogic {
	return &InsideOperationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsideInstallPlan Rpc End
func (l *InsideOperationAddLogic) InsideOperationAdd(in *intranetclient.AddInsideOperationReq) (*intranetclient.InsideOperationCommonResp, error) {
	var tmp model.InsideOperation
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.InsideOperationModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideOperationCommonResp{}, nil
}
