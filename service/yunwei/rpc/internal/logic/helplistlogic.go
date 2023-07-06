package logic

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"ywadmin-v3/common/xcmd"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelpListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHelpListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelpListLogic {
	return &HelpListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ConfigFileDelivery Rpc End
func (l *HelpListLogic) HelpList(in *yunweiclient.HelpListReq) (*yunweiclient.HelpListResp, error) {

	if in.GameName == "" {
		in.GameName = "all"
	}
	scriptFile := fmt.Sprintf("sh %s/template.sh -g %s", filepath.Dir(l.svcCtx.Config.Scripts.MaintainFilePath), in.GameName)
	job := xcmd.NewCommandJob(2*time.Minute, scriptFile)
	if !job.IsOk {
		return nil, xerr.NewErrMsg("生成模板文件失败" + job.ErrMsg)
	}

	return &yunweiclient.HelpListResp{
		Rows: ReadContentReturnMap(job.OutMsg),
	}, nil
}

func ReadContentReturnMap(content string) []*yunweiclient.HelpListData {

	templateData := make([]*yunweiclient.HelpListData, 0)
	for _, line := range strings.Split(content, "############") {
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			tmp := new(yunweiclient.HelpListData)
			compile := regexp.MustCompile(`操作:(.*)\(.*`)
			findString := compile.FindAllStringSubmatch(line, -1)
			if len(findString) > 0 {
				//ct := strings.Split(findString[0][1], "|")
				//if len(ct) > 1 {
				//	for _, v := range ct {
				//		tmp.Key = v
				//		tmp.Value = line
				//	}
				//} else {
				keys := strings.ReplaceAll(findString[0][1], "|", "或")
				tmp.Key = keys
				tmp.Value = line
				//}

			} else {
				compile = regexp.MustCompile(`操作:(.*)`)
				findString = compile.FindAllStringSubmatch(line, -1)
				if len(findString) > 0 {
					tmp.Key = findString[0][1]
					tmp.Value = line
				} else {

					tmp.Key = "其他"
					tmp.Value = line
				}
			}

			templateData = append(templateData, tmp)
		}

	}
	return templateData
}
