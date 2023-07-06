package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StgroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStgroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupListLogic {
	return &StgroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StgroupListLogic) StgroupList(in *adminclient.StgroupListReq) (*adminclient.StgroupListResp, error) {
	var (
		all *[]model.SysStgroup
		err error
	)
	filters := make([]interface{}, 0)
	filters = append(filters, "id__=", in.Id)

	if in.Current == 0 && in.PageSize == 0 {
		all, err = l.svcCtx.StgroupModel.FindAll(l.ctx)
	} else {
		all, err = l.svcCtx.StgroupModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	}

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询策略组列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	count, _ := l.svcCtx.StgroupModel.Count(l.ctx, filters...)

	var list []*adminclient.StgroupListData
	for _, info := range *all {
		list = append(list, &adminclient.StgroupListData{
			Id:             info.Id,
			StName:         info.StName,
			StJson:         info.StJson,
			StRemark:       info.StRemark,
			CreateBy:       info.CreateBy,
			CreateTime:     info.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   info.LastUpdateBy,
			LastUpdateTime: info.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        info.DelFlag,
		})
	}
	return &adminclient.StgroupListResp{
		Total: count,
		List:  list,
	}, nil
}
