package log

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogListLogic) SysLogList(req *types.ListSysLogReq) (*types.ListSysLogResp, error) {
	resp, err := l.svcCtx.AdminRpc.SysLogList(l.ctx, &admin.SysLogListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		DateRange: req.DateRange,

		UserName: req.UserName,
		Ip:       req.Ip,
	})
	if err != nil {
		return nil, err
	}
	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "IP",
			Value: "ip",
			Types: "input",
		},
		{
			Label: "用户名",
			Value: "userName",
			Types: "input",
		},
		{
			Label: "创建时间",
			Value: "dateRange",
			Types: "daterange",
		},
	}
	list := make([]*types.ListSysLogData, 0)

	for _, log := range resp.List {
		list = append(list, &types.ListSysLogData{
			Id:         log.Id,
			UserName:   log.UserName,
			Operation:  log.Operation,
			Method:     log.Method,
			Params:     log.Params,
			Time:       log.Time,
			Ip:         log.Ip,
			CreateTime: log.CreateTime,
		})
	}

	return &types.ListSysLogResp{
		Rows:   list,
		Total:  resp.Total,
		Filter: filterList,
	}, nil
}
