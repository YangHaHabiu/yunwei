package stgroup

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StgroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStgroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupListLogic {
	return &StgroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StgroupListLogic) StgroupList(req *types.ListStgroupReq) (*types.ListStgroupResp, error) {
	resp, err := l.svcCtx.AdminRpc.StgroupList(l.ctx, &admin.StgroupListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		StName:   req.StName,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListStgroupData, 0)
	for _, tmp := range resp.List {
		list = append(list, &types.ListStgroupData{
			Id:             tmp.Id,
			StName:         tmp.StName,
			StJson:         tmp.StJson,
			StRemark:       tmp.StRemark,
			CreateBy:       tmp.CreateBy,
			CreateTime:     tmp.CreateTime,
			LastUpdateBy:   tmp.LastUpdateBy,
			LastUpdateTime: tmp.LastUpdateTime,
		})
	}

	return &types.ListStgroupResp{
		Rows:  list,
		Total: resp.Total,
	}, nil
}
