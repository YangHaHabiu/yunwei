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

type InsideHostInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideHostInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideHostInfoListLogic {
	return &InsideHostInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideHostInfoListLogic) InsideHostInfoList(in *intranetclient.ListInsideHostInfoReq) (*intranetclient.ListInsideHostInfoResp, error) {
	var (
		count int64
		list  []*intranetclient.ListInsideHostInfoData
		all   *[]model.InsideHostInfo
		err   error
	)

	filters := make([]interface{}, 0)

	count, _ = l.svcCtx.InsideHostInfoModel.Count(l.ctx, filters...)
	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.InsideHostInfoModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.InsideHostInfoModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
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
	return &intranetclient.ListInsideHostInfoResp{
		Rows:  list,
		Total: count,
	}, nil
}
