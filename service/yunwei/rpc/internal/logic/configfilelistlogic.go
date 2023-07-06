package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileListLogic {
	return &ConfigFileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigFileListLogic) ConfigFileList(in *yunweiclient.ListConfigFileReq) (*yunweiclient.ListConfigFileResp, error) {
	var (
		count int64
		list  []*yunweiclient.ListConfigFileData
		all   *[]model.ConfigFileList
		err   error
	)

	filters := make([]interface{}, 0)

	count, _ = l.svcCtx.ConfigFileModel.Count(l.ctx, filters...)

	all, err = l.svcCtx.ConfigFileModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ListConfigFileResp{
		Rows:  list,
		Total: count,
	}, nil
}
