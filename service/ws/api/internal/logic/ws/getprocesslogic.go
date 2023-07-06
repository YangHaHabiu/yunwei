package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"ywadmin-v3/service/ws/api/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	// 允许等待的写入时间
	//writeWait = 30 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWaitx = 10 * time.Hour

	// Send pings to peer with this period. Must be less than pongWait.
	//pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// 最大的连接ID，每次连接都加1 处理
var (
	maxConnId int64
	taskIdx   int64
)

// 客户消息结构体
type clientMsg struct {
	clientId  int64
	msg       string
	tmps      string //临时消息
	taskIdTmp int64
}

// 停止结构体
type stopMsg struct {
	clientId int64
	taskId   int64
}

// 客户端读写消息
type wsMessage struct {
	// websocket.TextMessage 消息类型
	messageType int
	data        []byte
}

func init() {
	wsConnAll = make(map[int64]*wsConnection)
}

// ws 的所有连接
// 用于广播
var wsConnAll map[int64]*wsConnection

//var upgraderX = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//	// 允许所有的CORS 跨域请求，正式环境可以关闭
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道,加锁处理
	isClosed  bool
	closeChan chan byte // 关闭通知
	id        int64
}

func WsHandler(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext) {
	ctx, _ := context.WithCancel(context.Background())
	//defer cancel()
	// 应答客户端告知升级连接为websocket
	wsSocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("升级为websocket失败", err.Error())
		return
	}
	maxConnId++
	// TODO 如果要控制连接数可以计算，wsConnAll长度
	// 连接数保持一定数量，超过的部分不提供服务
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		id:        maxConnId,
	}
	wsConnAll[maxConnId] = wsConn
	//log.Println("当前在线人数", len(wsConnAll))
	cObj := clientMsg{
		clientId: wsConn.id,
	}
	chs := make(chan clientMsg)
	stopch := make(chan stopMsg)
	// 处理器,发送定时信息，避免意外关闭
	go wsConn.processLoop(chs, stopch, cObj, ctx, svcCtx)
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

type taskStruct struct {
	TaskId int64 `json:"taskId"`
}

type realTimeLog struct {
	LogTypes string `json:"logTypes"`
}

// 前端的时间线进度条样式结构
type TimeLineTree struct {
	Content string `json:"content"`
	Size    string `json:"size"`
	Type    string `json:"type,omitempty"`
}

// 处理队列中的消息
func (wsConn *wsConnection) processLoop(chs chan clientMsg, stopch chan stopMsg, cObj clientMsg, ctx context.Context, svcCtx *svc.ServiceContext) {
	// 处理消息队列中的消息
	// 获取到消息队列中的消息，处理完成后，发送消息给客户端
	for {

		msg, err := wsConn.wsRead()
		if err != nil {
			//logx.Error("获取消息出现错误", err.Error())
			break
		}

		log.Println("接收用户ID：[", wsConn.id, "]到消息为：", string(msg.data))

		//查询根据用户发送的消息进行查询对比
		var t *taskStruct
		err = json.Unmarshal(msg.data, &t)
		if err != nil {
			continue
		}

		if cObj.taskIdTmp != 0 {
			stopch <- stopMsg{
				clientId: wsConn.id,
				taskId:   cObj.taskIdTmp,
			}
		}

		go diffFun(t.TaskId, wsConn.id, chs, stopch, cObj, ctx, svcCtx)
		// 异步发送消息
		go func() {

			for {
				select {
				case ch1, ok := <-chs:
					if ok {
						// 修改以下内容把客户端传递的消息传递给处理程序
						err = wsConn.wsWrite(msg.messageType, []byte(ch1.msg+"\n"))
						if err != nil {
							fmt.Println(err)
							//logx.Error("发送消息给客户端出现错误", err.Error())
							break
						}
						cObj = ch1
					}
				}
			}
		}()
	}
}

// 处理消息队列中的消息
func (wsConn *wsConnection) wsReadLoop() {
	// 设置消息的最大长度
	wsConn.wsSocket.SetReadLimit(maxMessageSize)
	wsConn.wsSocket.SetReadDeadline(time.Now().Add(pongWaitx))
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			log.Println("消息读取出现错误", err)
			wsConn.close()
			return
		}
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列,消息入栈
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			return
		}
	}
}

// 发送消息给客户端
func (wsConn *wsConnection) wsWriteLoop() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				log.Println("发送消息给客户端发生错误", err)
				// 切断服务
				wsConn.close()
				return
			}
		case <-wsConn.closeChan:
			// 获取到关闭通知
			return
		case <-ticker.C:
			// 出现超时情况
			wsConn.wsSocket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := wsConn.wsSocket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 写入消息到队列中
func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("连接已经关闭")
	}
	return nil
}

// 读取消息队列中的消息
func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		// 获取到消息队列中的消息
		return msg, nil
	case <-wsConn.closeChan:

	}
	return nil, errors.New("连接已经关闭")
}

// 关闭连接
func (wsConn *wsConnection) close() {
	log.Println("关闭连接被调用了")
	wsConn.wsSocket.Close()
	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if wsConn.isClosed == false {
		wsConn.isClosed = true
		// 删除这个连接的变量
		delete(wsConnAll, wsConn.id)
		close(wsConn.closeChan)
	}
}

func diffFun(taskId, clientId int64, chs chan clientMsg, stopch chan stopMsg, cObj clientMsg, ctx context.Context, svcCtx *svc.ServiceContext) {
	for {
		select {
		case v, ok := <-stopch:
			if ok {
				if v.clientId == clientId && v.taskId == taskId {
					chs <- cObj
					return
				}
			}
		default:

		}
		get, err := svcCtx.YunWeiRpc.TasksGet(ctx, &yunweiclient.GetTasksReq{Id: taskId})
		if err != nil {
			//fmt.Println(err)
			logx.Error(err)
			return
		}
		rows := get.Rows
		vueTree := make([]*TimeLineTree, 0)
		for _, v := range rows {
			if v.Pid == 0 {
				continue
			}
			if v.Level == 3 && v.TaskStatus != 5 && v.TaskStatus != 6 {
				oneVue := new(TimeLineTree)
				//entity := &updateModel.Entity{Id: v.Pid}
				entity, err := svcCtx.YunWeiRpc.TasksGetOneById(ctx, &yunweiclient.GetTasksReq{Id: v.Pid})
				if err != nil {
					logx.Error(err)
					return
				}
				if entity == nil {
					logx.Error("查询任务对象为空，请重试")
					return
				}
				lables := fmt.Sprintf("%s-%s", entity.Name, v.Name)
				oneVue.Content = lables
				if v.TaskStatus == 3 {
					oneVue.Type = "success"
				}
				if v.TaskStatus == 2 {
					oneVue.Type = "danger"
				}
				oneVue.Size = "large"
				vueTree = append(vueTree, oneVue)
			}

		}

		marshal, _ := json.Marshal(vueTree)
		if cObj.tmps == "" || cObj.tmps != string(marshal) {
			cObj.tmps = string(marshal)
			cObj.taskIdTmp = taskId
			cObj.msg = string(marshal)
			chs <- cObj
		}

		time.Sleep(1 * time.Second)

	}
}
