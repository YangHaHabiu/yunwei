package inws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
	"io/ioutil"
	"net/http"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/service/inws/api/internal/svc"
	"ywadmin-v3/service/inws/api/internal/types"
)

const (
	writeWait  = 300 * time.Second
	pongWait   = 1 * time.Hour
	pingPeriod = (pongWait * 9) / 10
)

type taskStruct struct {
	TasksId int64 `json:"tasksId"`
}

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

func ServeWs(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext, req *types.GetInsideTasksLogReq) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			fmt.Println(err)
			return
		}
		return
	}
	defer ws.Close()
	_, err = tool.ParseToken(req.Token, svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		ws.WriteControl(websocket.CloseMessage,
			[]byte("解析token失败"), time.Now().Add(time.Second))
		return
	}

	//filename = fmt.Sprintf("%s/%d.txt", svcCtx.Config.ScriptsFilePath, req.TaskId)
	//content, _ := ioutil.ReadFile(filename)
	//ws.WriteMessage(websocket.TextMessage, content)
	ctx, canf := context.WithCancel(context.Background())
	//go writer(ws, filename, ctx)
	reader(ws, svcCtx, ctx, canf)

}

func reader(ws *websocket.Conn, svcCtx *svc.ServiceContext, ctx context.Context, canf context.CancelFunc) {
	defer ws.Close()
	ws.SetReadLimit(8192)
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
			var t *taskStruct
			err = json.Unmarshal(p, &t)
			if err != nil || t.TasksId == 0 {
				continue
			}
			//one, err := svcCtx.IntranetRpc.InsideTasksList(ctx, &intranetclient.ListInsideTasksReq{
			//	Current: 0, PageSize: 0, TasksId: t.TaskId,
			//})
			//if err != nil || len(one.Rows) != 1 {
			//	continue
			//}
			//fmt.Println(one.Rows[0].ProjectEn)

			filename = fmt.Sprintf("%s/%d.txt", svcCtx.Config.ScriptsFilePath, t.TasksId)
			if !tool.IsExist(filename) {
				continue
			}
			//杀掉子携程
			canf()
			//建立新的机制
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
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
