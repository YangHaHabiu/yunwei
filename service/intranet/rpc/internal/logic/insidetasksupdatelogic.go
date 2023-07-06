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

type InsideTasksUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksUpdateLogic {
	return &InsideTasksUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksUpdateLogic) InsideTasksUpdate(in *intranetclient.UpdateInsideTasksReq) (*intranetclient.InsideTasksCommonResp, error) {
	var tmp model.InsideTasks
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝更新数据失败，原因：" + err.Error())
	}
	err = l.svcCtx.InsideTasksModel.Update(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideTasksCommonResp{}, nil
}
