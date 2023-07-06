package file

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10 MB

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewFileUploadLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload() (resp *types.FileUploadResp, err error) {
	var fullPath string

	l.r.ParseMultipartForm(maxFileSize)
	file, handler, err := l.r.FormFile("myFile")
	if err != nil {
		return nil, xerr.NewErrMsg("上传文件失败，失败原因：" + err.Error())
	}
	action := l.r.FormValue("action")
	getwd, _ := os.Getwd()
	if action == "import_sql_files" {
		fullPath = filepath.Join(l.svcCtx.Config.Scripts.MaintainFilePath, "format_files", action)
	} else if action != "" {
		fullPath = filepath.Join(getwd, l.svcCtx.Config.Path, action)
	} else {
		fullPath = filepath.Join(getwd, l.svcCtx.Config.Path)
	}

	defer file.Close()
	if handler.Size > maxFileSize {
		return nil, xerr.NewErrMsg("上传文件失败，文件大小超过10M")
	}

	logx.Errorf("upload file: %+v, file size: %d, MIME header: %+v",
		handler.Filename, handler.Size, handler.Header)
	fileName := fmt.Sprintf("%d-%s", time.Now().UnixMicro(), handler.Filename)
	newPathFile := path.Join(fullPath, fileName)
	err = tool.CreateMutiDir(fullPath)
	if err != nil {
		return nil, xerr.NewErrMsg("上传目录创建失败，失败原因：" + err.Error())
	}

	tempFile, err := os.Create(newPathFile)
	if err != nil {
		return nil, xerr.NewErrMsg("上传文件失败，失败原因：" + err.Error())
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)
	//判断为上传的key，赋予权限
	if action == "upload_keys" {
		os.Chmod(newPathFile, 0600)
	}
	return &types.FileUploadResp{
		File:     fileName,
		FilePath: fullPath,
	}, nil
}
