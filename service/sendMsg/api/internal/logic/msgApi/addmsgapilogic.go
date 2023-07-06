package msgApi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"strings"
	"time"
	"ywadmin-v3/common/myip"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/sendMsg/model"

	"ywadmin-v3/service/sendMsg/api/internal/svc"
	"ywadmin-v3/service/sendMsg/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMsgApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAddMsgApiLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AddMsgApiLogic {
	return &AddMsgApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

var msgTypeMap = map[string]bool{
	"wechat":   true,
	"email":    true,
	"dingding": true,
	"feishu":   true,
}
var sendTypeMap = map[string]string{
	"1": "业务",
	"2": "合服",
	"3": "其他",
}

func (l *AddMsgApiLogic) AddMsgApi(req *types.AddMsgApiReq) error {
	//检测appkey是否正确
	var (
		appKeyIndex = -1
		flag        bool
		errmsg      string
	)
	for i, v := range l.svcCtx.Config.ApiAuthKey.List {
		if v.AppKey == req.AppKey {
			appKeyIndex = i
		}
	}
	if appKeyIndex == -1 {
		return xerr.NewErrMsg("appKey不正确，请检查")
	}
	appKey := l.svcCtx.Config.ApiAuthKey.List[appKeyIndex].AppKey
	appSecret := l.svcCtx.Config.ApiAuthKey.List[appKeyIndex].AppSecret
	now := time.Now().Unix()
	//加密校验
	if Sign(req.Ts, req.Sn, appKey, appSecret) == false {
		return xerr.NewErrMsg("校验加密失败")
	}
	//判断超时(10分钟超时)
	if (now-gconv.Int64(req.Ts)) > 600 || (now-gconv.Int64(req.Ts)) < -60 {
		return xerr.NewErrMsg("加密字符串已超10分钟，请重新生成加密字符")
	}
	//检测消息类型
	for _, v := range strings.Split(req.MsgType, ",") {
		if _, ok := msgTypeMap[v]; !ok {
			flag = true
			errmsg += v + ","
		}
	}
	if flag {
		return xerr.NewErrMsg(errmsg + "消息类型错误，请检查")
	}
	if _, ok := sendTypeMap[req.SendType]; !ok {
		return xerr.NewErrMsg("sendType发送类型错误，请检查")
	}

	for _, v := range strings.Split(req.MsgType, ",") {
		list := make([]string, 0)
		all, err := l.svcCtx.SendUserModel.FindAllByNames(l.ctx, req.MsgTo, v)
		if err != nil || len(*all) == 0 {
			return xerr.NewErrMsg("查询用户信息失败")
		}
		for _, v1 := range *all {
			list = append(list, v1.Result)
		}
		duplicate := tool.RemoveDuplicate(list)
		msg := new(model.SendMsgRecord)
		msg.SendType = req.SendType
		msg.MsgType = v
		msg.Status = "1"
		msg.MsgContent = strings.ReplaceAll(req.MsgContent, " ", "\n")
		msg.MsgTitle = req.MsgTitle
		msg.MsgTo = strings.Join(duplicate, ",")
		msg.CreateDate = time.Now().Unix()
		msg.AppKey = appKey
		msg.MsgLevel = req.MsgLevel
		msg.AccessIp = myip.GetCurrentIP(l.r)

		_, err = l.svcCtx.SendMsgRecordModel.Insert(l.ctx, msg)
		if err != nil {
			return xerr.NewErrMsg("写入消息队列失败")
		}
	}
	return nil
}

//校验sign
func Sign(ts, sign, appKey, secretCode string) bool {
	b := bytes.Buffer{}
	b.WriteString("appKey=")
	b.WriteString(appKey)
	b.WriteString("&appSecret=")
	b.WriteString(secretCode)
	b.WriteString("&ts=")
	b.WriteString(ts)
	fmt.Println(b.String())
	fmt.Println("md5:", gmd5.MustEncryptString(b.String()))
	if gmd5.MustEncryptString(b.String()) == sign {
		return true
	}
	return false
}
