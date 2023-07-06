package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotLogHistoryGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotLogHistoryGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotLogHistoryGetLogic {
	return &HotLogHistoryGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HotLogHistoryGetLogic) HotLogHistoryGet(in *yunweiclient.GetHotLogHistoryReq) (*yunweiclient.ListHotLogHistoryData, error) {
	one, err := l.svcCtx.HotLogHistoryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败" + err.Error())
	}
	var tmp yunweiclient.ListHotLogHistoryData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
