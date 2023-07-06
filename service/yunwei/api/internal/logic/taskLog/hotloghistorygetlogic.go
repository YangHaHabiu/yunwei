package taskLog

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotLogHistoryGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotLogHistoryGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotLogHistoryGetLogic {
	return &HotLogHistoryGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotLogHistoryGetLogic) HotLogHistoryGet(req *types.GetHotLogHistoryReq) (resp *types.ListHotLogHistoryData, err error) {
	get, err := l.svcCtx.YunWeiRpc.HotLogHistoryGet(l.ctx, &yunweiclient.GetHotLogHistoryReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var tmp types.ListHotLogHistoryData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
