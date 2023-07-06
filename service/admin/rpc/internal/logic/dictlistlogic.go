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

type DictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictListLogic {
	return &DictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictListLogic) DictList(in *adminclient.DictListReq) (*adminclient.DictListResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "types__=", in.Types)
	filters = append(filters, "id__=", in.Id)
	if in.Pid != 0 {
		if in.Pid == -2 {
			in.Pid = -1
			filters = append(filters, "pid__!=", in.Pid)
		} else {
			filters = append(filters, "pid__=", in.Pid)
		}

	}

	var (
		all   *[]model.SysDict
		err   error
		count int64
	)
	if in.Current == 0 && in.PageSize == 0 {
		all, err = l.svcCtx.DictModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.DictModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	}
	count, _ = l.svcCtx.DictModel.Count(l.ctx, filters...)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询字典列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrCode(xerr.ADMIN_DICTSELECT_ERROR)
	}
	var list []*adminclient.DictListData
	for _, dict := range *all {
		list = append(list, &adminclient.DictListData{
			Id:          dict.Id,
			Value:       dict.Value,
			Label:       dict.Label,
			Types:       dict.Types,
			Pid:         dict.Pid,
			Description: dict.Description,
			Sort:        dict.Sort,
		})
	}

	return &adminclient.DictListResp{
		Total: count,
		List:  list,
	}, nil
}
