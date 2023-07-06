package ugroup

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUgroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupListLogic {
	return &UgroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UgroupListLogic) UgroupList(req *types.ListUgroupReq) (*types.ListUgroupResp, error) {
	resp, err := l.svcCtx.AdminRpc.UgroupList(l.ctx, &admin.UgroupListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UgName:   req.UgName,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListUgroupData, 0)
	for _, ugroup := range resp.List {
		list = append(list, &types.ListUgroupData{
			Id:             ugroup.Id,
			UgName:         ugroup.UgName,
			UgJson:         ugroup.UgJson,
			CreateBy:       ugroup.CreateBy,
			CreateTime:     ugroup.CreateTime,
			LastUpdateBy:   ugroup.LastUpdateBy,
			LastUpdateTime: ugroup.LastUpdateTime,
		})
	}

	return &types.ListUgroupResp{

		Rows: list,

		Total: resp.Total,
	}, nil
}
