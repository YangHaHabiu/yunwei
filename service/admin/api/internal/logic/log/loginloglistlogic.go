package log

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogListLogic) LoginLogList(req *types.ListLoginLogReq) (*types.ListLoginLogResp, error) {
	resp, err := l.svcCtx.AdminRpc.LoginLogList(l.ctx, &admin.LoginLogListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		Status:    req.Status,
		DateRange: req.DateRange,
		Ip:        req.Ip,
		UserName:  req.UserName,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询登录日志列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListLoginLogData, 0)

	for _, log := range resp.List {
		list = append(list, &types.ListLoginLogData{
			Id:         log.Id,
			UserName:   log.UserName,
			Status:     log.Status,
			Ip:         log.Ip,
			CreateTime: log.CreateTime,
		})
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
		{
			Label: "状态",
			Value: "status",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "在线",
					Value: "online",
				},
				{
					Label: "退出",
					Value: "logout",
				},
				{
					Label: "离线",
					Value: "login",
				},
			},
		},
	}

	return &types.ListLoginLogResp{
		Rows:   list,
		Total:  resp.Total,
		Filter: filterList,
	}, nil
}
