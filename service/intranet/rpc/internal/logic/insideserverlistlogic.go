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

type InsideServerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideServerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideServerListLogic {
	return &InsideServerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideServerListLogic) InsideServerList(in *intranetclient.ListInsideServerReq) (*intranetclient.ListInsideServerResp, error) {
	var (
		count int64
		list  []*intranetclient.ListInsideServerData
		all   *[]model.InsideServerNew
		err   error
	)

	filters := make([]interface{}, 0)
	filters = append(filters,
		"project_id__=", in.ProjectId,
		"cluster_id__=", in.ClusterId,
		"build_type__=", in.BuildType,
	)
	if in.FeatureType == "all" {
		filters = append(filters,
			"feature_type__!=", "cdn",
		)
	} else {
		filters = append(filters,
			"feature_type__=", in.FeatureType,
		)
	}
	count, _ = l.svcCtx.InsideServerModel.Count(l.ctx, filters...)
	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.InsideServerModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.InsideServerModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
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
	return &intranetclient.ListInsideServerResp{
		Rows:  list,
		Total: count,
	}, nil
}
