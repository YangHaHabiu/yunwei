package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLabelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelListLogic {
	return &LabelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LabelListLogic) LabelList(in *adminclient.LabelListReq) (*adminclient.LabelListResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "view_label_type__in", in.LabelType,
		"view_label_values__=", in.LabelValues,
		"view_label_name__like", in.LabelName,
		"view_label_id__=", in.LabelId)

	var (
		count int64
		err   error
		all   *[]model.LabelView
	)
	if in.Current == 0 && in.PageSize == 0 {
		all, err = l.svcCtx.LabelModel.FindAll(l.ctx, filters...)
	} else {
		count, _ = l.svcCtx.LabelModel.Count(l.ctx, filters...)
		all, err = l.svcCtx.LabelModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	}

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrCode(xerr.ADMIN_LABELSELECT_ERROR)
	}

	var list []*adminclient.LabelListData
	for _, data := range *all {
		list = append(list, &adminclient.LabelListData{
			ViewLabelId:     data.ViewLabelId,
			ViewLabelName:   data.ViewLabelName,
			ViewLabelValues: data.ViewLabelValues,
			ViewLabelType:   data.ViewLabelType,
			ViewLabelRemark: data.ViewLabelRemark,
			ViewStopStatus:  data.ViewStopStatus,
		})
	}

	return &adminclient.LabelListResp{
		Total: count,
		List:  list,
	}, nil
}
