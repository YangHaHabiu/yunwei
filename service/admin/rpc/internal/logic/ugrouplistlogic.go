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

type UgroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUgroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupListLogic {
	return &UgroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UgroupListLogic) UgroupList(in *adminclient.UgroupListReq) (*adminclient.UgroupListResp, error) {
	var (
		count int64
		err   error
		list  []*adminclient.UgroupListData
		all   *[]model.SysUgroup
	)
	count, _ = l.svcCtx.UgroupModel.Count(l.ctx)
	if in.Current != 0 && in.PageSize != 0 {
		all, err = l.svcCtx.UgroupModel.FindPageListByPage(l.ctx, in.Current, in.PageSize)
	} else {
		all, err = l.svcCtx.UgroupModel.FindAll(l.ctx)
	}
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询用户组列表失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询用户组列表失败")
	}
	for _, info := range *all {
		list = append(list, &adminclient.UgroupListData{
			Id:             info.Id,
			UgName:         info.Ugname,
			UgJson:         info.UgJson,
			CreateBy:       info.CreateBy,
			CreateTime:     info.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   info.LastUpdateBy,
			LastUpdateTime: info.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        info.DelFlag,
		})
	}

	return &adminclient.UgroupListResp{
		Total: count,
		List:  list,
	}, nil
}
