package ws

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/webssh"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/ws/api/internal/svc"
	"ywadmin-v3/service/ws/api/internal/types"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

type WebSSHConfig struct {
	Record     bool
	RecPath    string
	RemoteAddr string
	User       string
	Password   string
	AuthModel  webssh.AuthModel
	PkPath     string
	JumpAddr   string
	SocksName  string
	SocksPwd   string
	Token      string
}

type WebSSH struct {
	*WebSSHConfig
}

func NewWebSSH(conf *WebSSHConfig) *WebSSH {
	return &WebSSH{
		WebSSHConfig: conf,
	}
}

var upgraderx = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSshHandler(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext, req *types.GetWebSshReq) {
	var (
		sshport        int64
		privateKeyPath string
		ipaddr         string
		socksName      string
		socksPwd       string
		index          int
	)
	if req.SshPort == 0 {
		sshport = 22
	} else {
		sshport = req.SshPort
	}

	if req.Cluster == "" || strings.Contains(req.Cluster, "CN") {
		index = 0
	} else {
		index = 1
	}
	privateKeyPath = svcCtx.Config.JumpServer[index].PrivateKeyPath
	ipaddr = svcCtx.Config.JumpServer[index].Ipaddr
	socksName = svcCtx.Config.JumpServer[index].SocksName
	socksPwd = svcCtx.Config.JumpServer[index].SocksPwd

	confing := &WebSSHConfig{
		Record:     true,
		RecPath:    "./rec/",
		RemoteAddr: fmt.Sprintf("%s:%d", req.Hostname, sshport),
		//User:       "root",
		PkPath:    privateKeyPath,
		AuthModel: webssh.PUBLICKEY,
		JumpAddr:  ipaddr,
		SocksName: socksName,
		SocksPwd:  socksPwd,
		Token:     req.Token,
	}
	handle := NewWebSSH(confing)
	handle.ServeConn(w, r, svcCtx)
}

func (wes WebSSH) ServeConn(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext) {
	wsConn, err := upgraderx.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer wsConn.Close()
	claims, err := tool.ParseToken(wes.Token, svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte("解析token失败"), time.Now().Add(time.Second))
		return
	}

	userIdStr := claims.Claims.(jwt.MapClaims)[ctxdata.CtxKeyJwtUserId]
	userId := gconv.Int64(userIdStr)
	ctx, cancel := context.WithCancel(context.Background())
	list, err := svcCtx.AdminRpc.UserList(ctx, &adminclient.UserListReq{
		UserId: userId,
	})
	if err != nil || len(list.List) != 1 {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte("获取用户id失败"), time.Now().Add(time.Second))
		return
	}
	userObj := list.List[0]
	var (
		sshUser    string
		sshKeyPath string
	)
	if strings.Contains(userObj.DeptName, "运维") {
		sshUser = "root"
		sshKeyPath = filepath.Join(wes.PkPath, "id_rsa")
	} else {
		sshUser = "tfgame"
		sshKeyPath = filepath.Join(wes.PkPath, "id_rsa_tfgame")
	}
	userName := claims.Claims.(jwt.MapClaims)[ctxdata.CtxKeyJwtUserName].(string)
	var config *webssh.SSHClientConfig

	switch wes.AuthModel {
	case webssh.PASSWORD:
		config = webssh.SSHClientConfigPassword(
			wes.RemoteAddr,
			sshUser,
			wes.Password,
			wes.JumpAddr,
			wes.SocksName,
			wes.SocksPwd,
		)
	case webssh.PUBLICKEY:
		config = webssh.SSHClientConfigPulicKey(
			wes.RemoteAddr,
			sshUser,
			sshKeyPath,
			wes.JumpAddr,
			wes.SocksName,
			wes.SocksPwd,
		)
	}

	client, err := webssh.NewSSHClient(config)
	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte(err.Error()), time.Now().Add(time.Second))
		return
	}

	defer client.Close()

	var recorder *webssh.Recorder
	if wes.Record {
		//mask := syscall.Umask(0)
		//defer syscall.Umask(mask)
		//新建目录
		split := strings.Split(wes.RemoteAddr, ":")
		recFullPath := fmt.Sprintf("%s/%s/%s/%s", wes.RecPath, userName, time.Now().Format("20060102"), split[0])
		if !tool.IsExist(recFullPath) {
			os.MkdirAll(recFullPath, 0766)
		}
		fileName := path.Join(recFullPath, fmt.Sprintf("%s_%s.cast", time.Now().Format("150405"), tool.RandCreator(10)))
		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0766)
		if err != nil {
			return
		}
		defer f.Close()
		recorder = webssh.NewRecorder(f)
	}

	turn, err := webssh.NewTurn(wsConn, client, recorder)

	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte(err.Error()), time.Now().Add(time.Second))
		return
	}
	defer turn.Close()

	var logBuff = bufPool.Get().(*bytes.Buffer)
	logBuff.Reset()
	defer bufPool.Put(logBuff)

	//ctx, cancel = context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = turn.LoopRead(logBuff, ctx)
		if err != nil {
			log.Printf("%#v", err)
		}
	}()
	go func() {
		defer wg.Done()
		err = turn.SessionWait()
		if err != nil {
			log.Printf("%#v", err)
		}
		cancel()
	}()
	wg.Wait()
}
