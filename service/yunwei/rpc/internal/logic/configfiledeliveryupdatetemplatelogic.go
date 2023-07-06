package logic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xcmd"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryUpdateTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileDeliveryUpdateTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryUpdateTemplateLogic {
	return &ConfigFileDeliveryUpdateTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigFileDeliveryUpdateTemplateLogic) ConfigFileDeliveryUpdateTemplate(in *yunweiclient.UpdateConfigFileDeliveryTemplateReq) (*yunweiclient.UpdateConfigFileDeliveryTemplateResp, error) {

	path := filepath.Join(l.svcCtx.Config.TemplateFilePath, in.ConfigName)
	//写入模版文件中
	err := tool.WriteFile(path, in.Content)
	if err != nil {
		return nil, xerr.NewErrMsg("写入文件模版失败，原因：" + err.Error())
	}

	//生成新的文件到文件中心
	configCenterPath := fmt.Sprintf("source /etc/profile;bash %s/auto_create_env.sh %s", l.svcCtx.Config.ConfigCenterPath, in.ProjectEn)
	job := xcmd.NewCommandJob(1*time.Minute, configCenterPath)
	if job.ErrMsg != "" || !job.IsOk {
		return nil, xerr.NewErrMsg("生成配置文件失败，原因：" + job.ErrMsg)
	}
	if job.IsTimeout {
		return nil, xerr.NewErrMsg("执行命令超时")
	}
	//更新数据库
	newFilePath := filepath.Join(l.svcCtx.Config.ConfigCenterPath, in.ProjectEn, in.ConfigName)
	err = CommonUpdateConfigFile(l.ctx, l.svcCtx, newFilePath, in.ConfigName, in.ProjectId)
	if err != nil {
		return nil, err
	}

	return &yunweiclient.UpdateConfigFileDeliveryTemplateResp{}, nil
}

//通过模版文件生成目标文件并更新数据库
func CommonCreateCenterFile(ctx context.Context, svcCtx *svc.ServiceContext, Content, ProjectEn, ConfigName string, ProjectId int64) error {
	list, err := svcCtx.ConfigMngLogModel.FindShPlatformInfoList(ctx, "view_project_en__=", ProjectEn)
	if err != nil {
		return xerr.NewErrMsg("查询配置文件需要平台信息失败，原因：" + err.Error())
	}
	ipAndRegionList, err := svcCtx.ConfigMngLogModel.FindShSplitIpAndRegionList(ctx, "view_project_en__=", ProjectEn)
	if err != nil {
		return xerr.NewErrMsg("查询配置文件需要区域信息失败，原因：" + err.Error())
	}
	platformInfoList := make([]string, 0)
	ipInfoList := make([]string, 0)
	regionInfoList := make([]string, 0)

	for _, v := range *list {
		platformInfoList = append(platformInfoList, "\t\t"+v.ViewPlatformInfo.String)
	}
	for _, v := range *ipAndRegionList {
		ipInfoList = append(ipInfoList, v.ViewSingleIpPool.String+"\n"+v.ViewCrossIpPool.String)
		regionInfoList = append(regionInfoList, "\t"+v.ViewSplitSingleRegion.String+"\n"+"\t"+v.ViewSplitCrossRegion.String)
	}
	newFileContent := fmt.Sprintf(Content, strings.Join(platformInfoList, "\n"), ProjectEn,
		strings.Join(ipInfoList, "\n"),
		strings.Join(regionInfoList, "\n"))
	filePaths := filepath.Join(svcCtx.Config.ConfigCenterPath, ProjectEn)
	if !tool.IsExist(filePaths) {
		tool.CreateMutiDir(filePaths)
	}

	newFilePath := filepath.Join(svcCtx.Config.ConfigCenterPath, ProjectEn, ConfigName)
	err = tool.WriteFile(newFilePath, newFileContent)
	if err != nil {
		return xerr.NewErrMsg("生成新文件失败，原因：" + err.Error())
	}
	err = CommonUpdateConfigFile(ctx, svcCtx, newFilePath, ConfigName, ProjectId)
	if err != nil {
		return xerr.NewErrMsg("更新新的配置文件失败，原因：" + err.Error())
	}
	return nil
}

//根据目标文件时间戳更新数据库
func CommonUpdateConfigFile(ctx context.Context, svcCtx *svc.ServiceContext, newFilePath, ConfigName string, ProjectId int64) error {
	oneObj, err := svcCtx.ConfigFileModel.FindOneByNameAndPrId(ctx, ProjectId, ConfigName)
	if err != nil {
		return err
	}
	stat, _ := os.Stat(newFilePath)
	lastModTime := stat.ModTime().Unix()
	if lastModTime != oneObj.FileModTime {
		oneObj.FileModTime = lastModTime
		err = svcCtx.ConfigFileModel.Update(ctx, oneObj)
		if err != nil {
			return err
		}
	}
	return nil
}
