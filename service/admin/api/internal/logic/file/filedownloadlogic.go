package file

import (
	"context"
	"io"
	"io/ioutil"
	"path/filepath"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer io.Writer
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer io.Writer) *FileDownloadLogic {
	return &FileDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *FileDownloadLogic) FileDownload(req *types.DownloadReq) error {
	logx.Errorf("download %s", req.File)
	var filePathX string
	if req.Action == "import_sql_files" {
		filePathX = filepath.Join(l.svcCtx.Config.Path, "tasks_download_files", req.File)
	} else {
		filePathX = filepath.Join(l.svcCtx.Config.Path, req.File)
	}
	body, err := ioutil.ReadFile(filePathX)
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
