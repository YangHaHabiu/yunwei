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

type InsideProxyHostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProxyHostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostListLogic {
	return &InsideProxyHostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideProxyHostListLogic) InsideProxyHostList(in *intranetclient.ListInsideProxyHostReq) (*intranetclient.ListInsideProxyHostResp, error) {
	var (
		count int64
		list  []*intranetclient.ListInsideProxyHostData
		all   *[]model.InsideProxyHostNew
		err   error
	)

	filters := make([]interface{}, 0)

	count, _ = l.svcCtx.InsideProxyHostModel.Count(l.ctx, filters...)
	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.InsideProxyHostModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.InsideProxyHostModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
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
	return &intranetclient.ListInsideProxyHostResp{
		Rows:  list,
		Total: count,
	}, nil
}
