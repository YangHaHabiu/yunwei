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

type InsideOperationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideOperationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideOperationListLogic {
	return &InsideOperationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideOperationListLogic) InsideOperationList(in *intranetclient.ListInsideOperationReq) (*intranetclient.ListInsideOperationResp, error) {
	var (
		count int64
		list  []*intranetclient.ListInsideOperationData
		all   *[]model.InsideOperationNew
		err   error
	)

	filters := make([]interface{}, 0)
	filters = append(filters, "project_id__=", in.ProjectId,
		"oper_type__=", in.OperType,
	)
	count, _ = l.svcCtx.InsideOperationModel.Count(l.ctx, filters...)
	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.InsideOperationModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.InsideOperationModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
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
	return &intranetclient.ListInsideOperationResp{
		Rows:  list,
		Total: count,
	}, nil
}
