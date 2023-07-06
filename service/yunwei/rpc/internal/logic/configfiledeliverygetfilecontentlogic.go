package logic

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryGetFileContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileDeliveryGetFileContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryGetFileContentLogic {
	return &ConfigFileDeliveryGetFileContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigFileDeliveryGetFileContentLogic) ConfigFileDeliveryGetFileContent(in *yunweiclient.ConfigFileDeliveryGetFileContentReq) (*yunweiclient.ConfigFileDeliveryGetFileContentResp, error) {
	var path string
	if in.Option == "template" {
		path = filepath.Join(l.svcCtx.Config.TemplateFilePath)
	} else {
		path = filepath.Join(l.svcCtx.Config.ConfigCenterPath, in.ProjectEn)
	}
	if !tool.IsExist(path) {
		tool.CreateMutiDir(path)
	}
	pathFile := filepath.Join(path, in.ConfigName)
	content, err := ioutil.ReadFile(pathFile)
	if err != nil {
		return nil, xerr.NewErrMsg("读取文件失败，原因：" + err.Error())
	}

	return &yunweiclient.ConfigFileDeliveryGetFileContentResp{
		Content: string(content),
	}, nil
}
