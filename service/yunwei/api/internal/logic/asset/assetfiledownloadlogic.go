package asset

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/common/xsshClient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetFileDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer io.Writer
}

func NewAssetFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext, w io.Writer) *AssetFileDownloadLogic {
	return &AssetFileDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: w,
	}
}

func (l *AssetFileDownloadLogic) AssetFileDownload(req *types.AssetFileDownloadReq) error {

	var port int
	if req.Port == 0 {
		port = 22
	} else {
		port = req.Port
	}
	sshClient, err := xsshClient.NewSSHClient(req.Hostname, port, xsshClient.AuthConfig{User: "root", KeyFile: "../../ws/api/key/id_rsa"})
	if err != nil {
		return err
	}
	defer sshClient.Close()
	localPath := path.Join(l.svcCtx.Config.Path, "download", time.Now().Format("20060102"))
	err = tool.CreateMutiDir(localPath)
	if err != nil {
		return xerr.NewErrMsg("下载目录创建失败，失败原因：" + err.Error())
	}
	localPath = path.Join(localPath, fmt.Sprintf("%d.%s", time.Now().Unix(), req.File))
	remotePath := path.Join(req.Path, req.File)
	_, err = sshClient.Download(remotePath, localPath)
	if err != nil {
		return err
	}
	logx.Errorf("download %s", req.File)

	body, err := ioutil.ReadFile(localPath)
	if err != nil {
		return xerr.NewErrMsg("文件不存在，失败原因：" + err.Error())
	}

	n, err := l.writer.Write(body)
	if err != nil {
		return xerr.NewErrMsg("下载文件失败，失败原因：" + err.Error())
	}

	if n < len(body) {
		return io.ErrClosedPipe
	}

	return nil
}
