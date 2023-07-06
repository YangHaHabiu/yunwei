package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotLogHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotLogHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotLogHistoryAddLogic {
	return &HotLogHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HotLogHistoryAddLogic) HotLogHistoryAdd(in *yunweiclient.AddHotLogHistoryReq) (*yunweiclient.AddHotLogHistoryResp, error) {
	var tmp model.HotLogHistory
	err := copier.Copy(&tmp, in.Data)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}

	_, err = l.svcCtx.HotLogHistoryModel.Insert(l.ctx, &tmp, in.Uid)
	if err != nil {
		return nil, xerr.NewErrMsg("新增数据失败，原因：" + err.Error())
	}

	return &yunweiclient.AddHotLogHistoryResp{}, nil
}
