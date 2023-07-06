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

type InsideTasksPidAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksPidAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksPidAddLogic {
	return &InsideTasksPidAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsideTasksLogs Rpc End
func (l *InsideTasksPidAddLogic) InsideTasksPidAdd(in *intranetclient.AddInsideTasksPidReq) (*intranetclient.InsideTasksPidCommonResp, error) {
	var tmp model.InsideTasksPid
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.InsideTasksPidModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideTasksPidCommonResp{}, nil
}
