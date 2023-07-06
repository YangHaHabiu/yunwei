package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/service/ws/api/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Minute
	// Time allowed to read the next pong message from the client.
	pongWait = 10 * time.Hour
	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Poll file for changes with this period.
	filePeriod = 1 * time.Second
)

var (
	filename string
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	stop chan struct{}
)

func ServeWs(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext) {
	ctx, canf := context.WithCancel(context.Background())
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			fmt.Println(err)
			return
		}
		return
	}
	defer ws.Close()

	reader(ws, svcCtx, ctx, "", canf)
}

func reader(ws *websocket.Conn, svcCtx *svc.ServiceContext, ctx context.Context, newPath string, canf context.CancelFunc) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			continue
		}
		if p == nil {
			continue
		}
		if string(p) == "exit" {
			ws.Close()
		}
		if string(p) != "" {
			if newPath == "" {
				var t *taskStruct
				err = json.Unmarshal(p, &t)
				if err != nil || t.TaskId == 0 {
					continue
				}
				list, err := svcCtx.YunWeiRpc.TasksList(ctx, &yunweiclient.ListTasksReq{
					Current: 0, PageSize: 0, Id: t.TaskId,
				})
				if err != nil || len(list.Rows) != 1 {
					continue
				}
				filename = fmt.Sprintf("%s/task_info/%s/%s/%d/main.log", filepath.Dir(svcCtx.Config.Scripts.MaintainFilePath), list.Rows[0].ProjectEn, list.Rows[0].Types, t.TaskId)
			} else if newPath == "realTimeLog" {
				var t *realTimeLog
				err = json.Unmarshal(p, &t)
				if err != nil || t.LogTypes == "" {
					continue
				}
				switch t.LogTypes {
				case "install":
					filename = fmt.Sprintf("%s/Log/once_run_main.log", filepath.Dir(svcCtx.Config.Scripts.InstallFilePath))
				case "combine":
					filename = fmt.Sprintf("%s/log/run_main.log", filepath.Dir(svcCtx.Config.Scripts.CombineFilePath))
				case "migrate":
					filename = fmt.Sprintf("%s/log/once_run_main.log", filepath.Dir(svcCtx.Config.Scripts.MigrateFilePath))
				}

			}
			if !tool.IsExist(filename) {
				continue
			}
			canf()
			ctx, canf = context.WithCancel(context.Background())

			//读取整个文件信息
			content, _ := ioutil.ReadFile(filename)
			ws.WriteMessage(websocket.TextMessage, content)
			go writer(ws, filename, ctx)
		}

	}
}

func tailFile(filename string) *tail.Tail {
	tailfs, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,                                 // 文件被移除或被打包，需要重新打开
		Follow:    true,                                 // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取。
		MustExist: false,                                // 如果文件不存在，是否推出程序，false是不退出
		Poll:      true,
	})

	if err != nil {
		return nil
	}
	return tailfs
}

func writer(ws *websocket.Conn, filename string, ctx context.Context) {
	tailfs := tailFile(filename)
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-tailfs.Lines:
			if ok {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Text+"\n")); err != nil {
					return
				}
			} else {
				return
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
