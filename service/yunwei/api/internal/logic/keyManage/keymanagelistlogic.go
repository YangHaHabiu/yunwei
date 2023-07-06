package keyManage

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeyManageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeyManageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeyManageListLogic {
	return &KeyManageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeyManageListLogic) KeyManageList(req *types.ListKeyManageReq) (resp *types.ListKeyManageResp, err error) {
	tmp := make([]*types.ListKeyManageData, 0)
	list, err := l.svcCtx.YunWeiRpc.KeyManageList(l.ctx, &yunweiclient.ListKeyManageReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	//自定义筛选条件

	return &types.ListKeyManageResp{
		Rows:  tmp,
		Total: list.Total,
		//Filter: filterList,
	}, nil
}
