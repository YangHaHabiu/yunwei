package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLogListLogic) SysLogList(in *adminclient.SysLogListReq) (*adminclient.SysLogListResp, error) {

	filters := make([]interface{}, 0)
	filters = append(filters,
		"create_time__range", in.DateRange,
		"user_name__=", in.UserName,
		"ip__=", in.Ip,
	)

	all, err := l.svcCtx.SysLogModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	count, _ := l.svcCtx.SysLogModel.Count(l.ctx, filters...)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询系统日志列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrCode(xerr.ADMIN_SYSLOGSELECT_ERROR)
	}
	var list []*adminclient.SysLogListData
	for _, log := range *all {
		list = append(list, &adminclient.SysLogListData{
			Id:         log.Id,
			UserName:   log.UserName,
			Operation:  log.Operation,
			Method:     log.Method,
			Params:     log.Params,
			Time:       log.Time,
			Ip:         log.Ip,
			CreateTime: log.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &adminclient.SysLogListResp{
		Total: count,
		List:  list,
	}, nil
}
