package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeatureServerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerListLogic {
	return &FeatureServerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeatureServerListLogic) FeatureServerList(in *yunweiclient.ListFeatureServerReq) (*yunweiclient.ListFeatureServerResp, error) {
	var (
		count int64
		list  []*yunweiclient.ListFeatureServerData
		all   *[]model.FeatureServerInfoNew
		err   error
	)
	filters := make([]interface{}, 0)
	filters = append(filters,
		"fsi.feature_server_id__=", in.FeatureServerId,
		"fsi.remark__like", in.Remark,
		"fsi.feature_server_info__regexp", in.Ip,
		"fsi.feature_server_info->'$.type'__in", in.Feature,
		"fsi.feature_server_info->'$.domain'__regexp", in.Domain,
		"p.project_id__in", in.ProjectIds,
	)

	count, _ = l.svcCtx.FeatureServerInfoModel.Count(l.ctx, filters...)

	all, err = l.svcCtx.FeatureServerInfoModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	for _, v := range *all {
		list = append(list, &yunweiclient.ListFeatureServerData{
			FeatureServerId:   v.FeatureServerId,
			ProjectId:         v.ProjectId,
			FeatureServerInfo: v.FeatureServerInfo,
			Remark:            v.Remark,
			ProjectCn:         v.ProjectCn,
			ProjectEn:         v.ProjectEn,
			DelFlag:           v.DelFlag,
		})
	}
	return &yunweiclient.ListFeatureServerResp{
		Rows:  list,
		Total: count,
	}, nil
}
