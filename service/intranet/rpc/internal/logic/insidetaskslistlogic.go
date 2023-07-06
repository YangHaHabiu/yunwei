package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/model"

	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/metadata"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksListLogic {
	return &InsideTasksListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksListLogic) InsideTasksList(in *intranetclient.ListInsideTasksReq) (*intranetclient.ListInsideTasksResp, error) {
	var (
		count int64
		list  []*intranetclient.ListInsideTasksData
		all   *[]model.InsideTasksNew
		err   error
	)

	filters := make([]interface{}, 0)
	filters = append(filters,
		"project_id__=", in.ProjectId,
		"project_id__in", in.ProjectIds,
		"cluster_id__=", in.ClusterId,
		"server_id__=", in.ServerId,
		"version_id__=", in.ServerId,
		"tasks_type__=", in.TasksType,
		"id__=", in.TasksId,
		"start_time__xrange", in.StartTime,
		"status__in", in.Status,
	)

	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.InsideTasksModel.FindAll(l.ctx, filters...)
	} else {
		if in.RecentSubmit != "" {
			var uid string
			if md, ok := metadata.FromIncomingContext(l.ctx); ok {
				uid = md.Get("uid")[0]
			}
			filters = append(filters, "create_by__=", uid)
			limits := gconv.Int64(in.RecentSubmit)
			if limits == 0 {
				limits = 5
			}
			all, err = l.svcCtx.InsideTasksModel.FindPageListByPage(l.ctx, 0, limits, filters...)
		} else {
			all, err = l.svcCtx.InsideTasksModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)

		}
	}
	count, _ = l.svcCtx.InsideTasksModel.Count(l.ctx, filters...)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &intranetclient.ListInsideTasksResp{
		Rows:  list,
		Total: count,
	}, nil
}
