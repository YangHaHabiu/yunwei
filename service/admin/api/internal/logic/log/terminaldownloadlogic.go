package log

import (
	"context"
	"io"
	"io/ioutil"
	"path"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TerminalDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer io.Writer
}

func NewTerminalDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer io.Writer) *TerminalDownloadLogic {
	return &TerminalDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *TerminalDownloadLogic) TerminalDownload(req *types.DetailTerminalLogReq) error {
	pathx := path.Join(l.svcCtx.Config.RecPath.FullPath, "/", req.File)
	body, err := ioutil.ReadFile(pathx)
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
