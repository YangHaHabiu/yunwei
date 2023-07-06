package asset

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"ywadmin-v3/common/xsshClient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetFileLogic {
	return &AssetFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetFileLogic) AssetFile(req *types.AssetFileReq) (resp *types.AssetFileResp, err error) {
	var port int
	if req.Port == 0 {
		port = 22
	} else {
		port = req.Port
	}
	sshClient, err := xsshClient.NewSSHClient(req.Hostname, port, xsshClient.AuthConfig{User: "root", KeyFile: "../../ws/api/key/id_rsa"})
	if err != nil {
		return nil, err
	}
	defer sshClient.Close()

	cmd := fmt.Sprintf("ls -lha %s --time-style=\"+%%Y-%%m-%%d %%H:%%I:%%S\"|sort  -k2nr |grep -v total|grep -v grep", req.Path)
	execinfo, err := sshClient.Exec(cmd)
	if err != nil {
		return nil, err
	}
	resp = new(types.AssetFileResp)
	resp.Rows = readLine(execinfo.OutputString(), req.Path)
	return
}

func readLine(content, pathx string) (tmp []*types.AssetFileData) {
	//允许访问的目录
	allPath := []string{
		"/data", "/root", "/tmp",
	}
	split := strings.Split(content, "\n")
	tmp = make([]*types.AssetFileData, 0)
	for _, v := range split {
		r := regexp.MustCompile(` -> .*`)
		v = r.ReplaceAllString(v, "")

		compile := regexp.MustCompile(`[ ]+`)
		x := compile.Split(v, -1)

		var (
			kind   string
			isLink bool
			size   string
		)

		if strings.HasPrefix(x[0], "p") {
			kind = "p"
		} else if strings.HasPrefix(x[0], "c") {
			kind = "c"
		} else if strings.HasPrefix(x[0], "d") {
			kind = "d"
		} else if strings.HasPrefix(x[0], "b") {
			kind = "b"
		} else if strings.HasPrefix(x[0], "_") {
			kind = "_"
		} else if strings.HasPrefix(x[0], "l") {
			kind = "l"
			isLink = true
		} else if strings.HasPrefix(x[0], "s") {
			kind = "s"
		} else {
			kind = "?"
		}
		if len(x) >= 7 {
			if x[len(x)-1] == "." || x[len(x)-1] == ".." {
				continue
			}
			mustCompile := regexp.MustCompile(`,`)
			allString := mustCompile.FindAllString(x[4], -1)
			if len(allString) == 0 {
				size = x[4]
			}

			for _, vx := range allPath {
				if strings.HasPrefix(pathx, vx) && !strings.Contains(pathx, ".") {
					tmp = append(tmp, &types.AssetFileData{
						IsLink: isLink,
						Kind:   kind,
						Name:   x[len(x)-1],
						Date:   x[len(x)-3] + " " + x[len(x)-2],
						Size:   size,
						Code:   x[0],
					})
				} else if pathx == "/" {
					if x[len(x)-1] == strings.ReplaceAll(vx, "/", "") {
						tmp = append(tmp, &types.AssetFileData{
							IsLink: isLink,
							Kind:   kind,
							Name:   x[len(x)-1],
							Date:   x[len(x)-3] + " " + x[len(x)-2],
							Size:   size,
							Code:   x[0],
						})
					}
				}

			}
		}

	}

	return
}
