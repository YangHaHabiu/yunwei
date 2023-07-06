package ws

import (
	"context"
	"net/http"
	"ywadmin-v3/service/ws/api/internal/svc"
	"ywadmin-v3/service/ws/api/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetInstallLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInstallLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInstallLogLogic {
	return &GetInstallLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInstallLogLogic) GetInstallLog(req *types.GetInstallLogReq, w http.ResponseWriter, r *http.Request) {
	//ctx, _ := context.WithCancel(context.Background())
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			l.Logger.Error("链接失败" + err.Error())
			return
		}
		return
	}

	defer ws.Close()
	// if req.FileName == "" || req.FileName == "ALL" {
	// 	filename = fmt.Sprintf("%s/Log/once_run_main.log", filepath.Dir(l.svcCtx.Config.Scripts.InstallFilePath))
	// } else {
	// 	filename = fmt.Sprintf("%s/Log/%s/%s", filepath.Dir(l.svcCtx.Config.Scripts.InstallFilePath), req.GameName, req.FileName)
	// }
	// content, _ := ioutil.ReadFile(filename)
	// ws.WriteMessage(websocket.TextMessage, content)
	ctx, canf := context.WithCancel(context.Background())
	//go writer(ws, filename, ctx)

	reader(ws, l.svcCtx, ctx, "realTimeLog", canf)
}
