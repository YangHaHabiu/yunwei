package strategy

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StrategyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StrategyListLogic {
	return &StrategyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StrategyListLogic) StrategyList(req *types.ListstrategyReq) (*types.ListstrategyResp, error) {
	resp, err := l.svcCtx.AdminRpc.StrategyList(l.ctx, &admin.StrategyListReq{
		StName:   req.StName,
		StLevel:  req.StLevel,
		StPid:    req.StPid,
		StIsAuth: 1,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListstrategyData, 0)
	for _, tmp := range resp.List {
		list = append(list, &types.ListstrategyData{
			Id:       tmp.Id,
			StName:   tmp.StName,
			StLevel:  tmp.StLevel,
			StPid:    tmp.StPid,
			StRemark: tmp.StRemark,
			StIsAuth: tmp.StIsAuth,
		})
	}

	return &types.ListstrategyResp{
		Rows: list,
	}, nil
}
