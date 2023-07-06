package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/model"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksLogsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksLogsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksLogsListLogic {
	return &InsideTasksLogsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksLogsListLogic) InsideTasksLogsList(in *intranetclient.ListInsideTasksLogsReq) (*intranetclient.ListInsideTasksLogsResp, error) {
	var (
		count int64
		list  []*intranetclient.ListInsideTasksLogsData
		all   *[]model.InsideTasksLogs
		err   error
	)

	filters := make([]interface{}, 0)

	count, _ = l.svcCtx.InsideTasksLogsModel.Count(l.ctx, filters...)
	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.InsideTasksLogsModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.InsideTasksLogsModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	}
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &intranetclient.ListInsideTasksLogsResp{
		Rows:  list,
		Total: count,
	}, nil
}
