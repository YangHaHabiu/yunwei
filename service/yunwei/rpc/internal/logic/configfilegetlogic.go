package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileGetLogic {
	return &ConfigFileGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigFileGetLogic) ConfigFileGet(in *yunweiclient.GetConfigFileReq) (*yunweiclient.ListConfigFileData, error) {
	one, err := l.svcCtx.ConfigFileModel.FindOne(l.ctx, in.ConfigFileId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp yunweiclient.ListConfigFileData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil

}
