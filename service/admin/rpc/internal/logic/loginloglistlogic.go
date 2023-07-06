package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogListLogic) LoginLogList(in *adminclient.LoginLogListReq) (*adminclient.LoginLogListResp, error) {

	var (
		count int64
		all   *[]model.SysLoginLog
		err   error
	)
	filters := make([]interface{}, 0)
	filters = append(filters,
		"create_time__range", in.DateRange,
		"status__regexp", in.Status,
		"user_name__=", in.UserName,
		"ip__=", in.Ip,
	)

	all, err = l.svcCtx.LoginLogModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	count, _ = l.svcCtx.LoginLogModel.Count(l.ctx, filters...)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询登录记录列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*adminclient.LoginLogListData
	for _, log := range *all {
		list = append(list, &adminclient.LoginLogListData{
			Id:         log.Id,
			UserName:   log.UserName,
			Status:     log.Status,
			Ip:         log.Ip,
			CreateTime: log.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &adminclient.LoginLogListResp{
		Total: count,
		List:  list,
	}, nil
}
