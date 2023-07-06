package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"os"
	"path/filepath"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileAddLogic {
	return &ConfigFileAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TaskLogHistroy Rpc End
func (l *ConfigFileAddLogic) ConfigFileAdd(in *yunweiclient.AddConfigFileReq) (*yunweiclient.ConfigFileCommonResp, error) {
	var tmp model.ConfigFile
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}

	one, err := l.svcCtx.AdminRpc.ProjectGetOne(l.ctx, &adminclient.ProjectGetOneReq{ProjectId: tmp.ProjectId})
	if err != nil {
		return nil, xerr.NewErrMsg("查询项目数据失败，原因：" + err.Error())
	}

	newFilePath := filepath.Join(l.svcCtx.Config.ConfigCenterPath, one.ProjectEn, in.One.Name)
	stat, _ := os.Stat(newFilePath)
	lastModTime := stat.ModTime().Unix()
	tmp.FileModTime = lastModTime
	_, err = l.svcCtx.ConfigFileModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}

	return &yunweiclient.ConfigFileCommonResp{}, nil
}
