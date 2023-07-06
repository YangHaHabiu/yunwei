package logic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xtime"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskGetInstallLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskGetInstallLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskGetInstallLogListLogic {
	return &TaskGetInstallLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskGetInstallLogListLogic) TaskGetInstallLogList(in *yunweiclient.ListInstallLogListReq) (*yunweiclient.ListInstallLogListResp, error) {

	files := make([]string, 0)
	result := make([]*yunweiclient.ListInstallLogListData, 0)
	newPath := fmt.Sprintf("%s/Log/%s/", filepath.Dir(l.svcCtx.Config.Scripts.InstallFilePath), in.GameName)
	files, _ = tool.GetAllFile(newPath, files)
	startTime, endTime := xtime.GetDateTime(time.Now().Format("2006-01-02"))
	for i := 0; i < len(files); i++ {
		finfo, _ := os.Stat(newPath + files[i])
		linuxFileAttr := finfo.ModTime()
		if linuxFileAttr.Unix() >= startTime && linuxFileAttr.Unix() <= endTime && strings.Contains(files[i], "install_game") {
			result = append(result, &yunweiclient.ListInstallLogListData{
				Name: files[i],
			})
		}
	}
	return &yunweiclient.ListInstallLogListResp{
		Data: result,
	}, nil
}
