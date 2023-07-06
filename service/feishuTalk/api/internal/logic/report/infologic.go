package report

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"ywadmin-v3/service/feishuTalk/api/internal/svc"
	"ywadmin-v3/service/feishuTalk/api/internal/types"
	"ywadmin-v3/service/feishuTalk/model"
	"ywadmin-v3/service/feishuTalk/utils/encrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(r *http.Request) (resp *types.ReportResp, err error) {

	out, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(out))
	var tmp model.Challenge
	err = json.Unmarshal(out, &tmp)
	if err != nil {
		return nil, err
	}
	resp = new(types.ReportResp)
	if tmp.Challenge != "" && tmp.Type == "url_verification" {

		resp.Challenge = tmp.Challenge
		resp.Token = tmp.Token
		resp.Type = tmp.Type
		return resp, nil
	}

	var tmp1 model.ChallengeEncrypt
	err = json.Unmarshal(out, &tmp1)
	if err != nil {
		return
	}
	if tmp1.Encrypt != "" {
		s, err := encrypt.Decrypt(tmp1.Encrypt, l.svcCtx.Config.EncryptKey)
		if err != nil {
			return nil, err
		}
		fmt.Println("解密后：", s)
		var tmp2 model.Challenge
		err = json.Unmarshal([]byte(s), &tmp2)
		if err != nil {
			return nil, err
		}
		resp.Challenge = tmp2.Challenge
		resp.Token = tmp2.Token
		resp.Type = tmp2.Type
		return resp, nil
	}

	return
}
