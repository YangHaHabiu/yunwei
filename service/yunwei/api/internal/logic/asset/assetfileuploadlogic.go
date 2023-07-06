package asset

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/common/xsshClient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 100 << 20 // 100 MB

type AssetFileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAssetFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AssetFileUploadLogic {
	return &AssetFileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AssetFileUploadLogic) AssetFileUpload(req *types.AssetFileReq) error {
	l.r.ParseMultipartForm(maxFileSize)
	file, handler, err := l.r.FormFile("file")
	if err != nil {
		return xerr.NewErrMsg("上传文件失败，失败原因：" + err.Error())
	}
	defer file.Close()
	if handler.Size > maxFileSize {
		return xerr.NewErrMsg("上传文件失败，文件大小超过100M")
	}

	logx.Errorf("upload file: %+v, file size: %d, MIME header: %+v",
		handler.Filename, handler.Size, handler.Header)
	localPath := path.Join(l.svcCtx.Config.Path, "upload", time.Now().Format("20060102"))
	filename := fmt.Sprintf("%d.%s", time.Now().Unix(), handler.Filename)
	pathx := path.Join(localPath, filename)
	err = tool.CreateMutiDir(localPath)
	if err != nil {
		return xerr.NewErrMsg("上传目录创建失败，失败原因：" + err.Error())
	}
	tempFile, err := os.Create(pathx)

	if err != nil {
		return xerr.NewErrMsg("上传文件失败，失败原因：" + err.Error())
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)
	var port int
	if req.Port == 0 {
		port = 22
	} else {
		port = req.Port
	}
	sshClient, err := xsshClient.NewSSHClient(req.Hostname, port, xsshClient.AuthConfig{User: "root", KeyFile: "../../ws/api/key/id_rsa"})
	if err != nil {
		logx.Errorf("链接目标机器ssh失败原因： err-->", err)
		return err
	}
	defer sshClient.Close()
	getwd, _ := os.Getwd()

	_, err = sshClient.Upload(pathx, req.Path+"/"+filename)
	if err != nil {
		logx.Errorf("localPath:%s destPath:%s hostname:%s 上传目标机器文件失败原因： err-->%s", getwd+"/"+pathx, req.Path, req.Hostname, err)
		return err
	}

	return nil
}
