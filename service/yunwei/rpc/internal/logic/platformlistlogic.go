package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlatformListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlatformListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformListLogic {
	return &PlatformListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlatformListLogic) PlatformList(in *yunweiclient.ListPlatformReq) (*yunweiclient.ListPlatformResp, error) {
	var (
		count int64
		list  []*yunweiclient.ListPlatformData
		all   *[]model.PlatformList
		err   error
	)

	filters := make([]interface{}, 0)
	filters = append(filters, "project_id__in", in.ProjectIds,
		"id__=", in.Id,
		"platform_en not__like", in.Types,
		//"del_flag__!=", in.DelFlag,
		"platform_ex__in", in.PlatformInfo,
		"label_names__like", in.Label,
	)
	if in.PlatformType != "all" {
		filters = append(filters,
			"del_flag__!=", in.DelFlag,
		)
	}

	count, _ = l.svcCtx.PlatformModel.Count(l.ctx, filters...)

	if in.Current == 0 && in.PageSize == 0 {
		all, err = l.svcCtx.PlatformModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.PlatformModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)

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
	return &yunweiclient.ListPlatformResp{
		Rows:  list,
		Total: count,
	}, nil

}
