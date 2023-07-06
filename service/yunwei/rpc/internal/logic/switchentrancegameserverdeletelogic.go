package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"
	"ywadmin-v3/common/xerr"

	"github.com/gogf/gf/util/gconv"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchEntranceGameserverDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSwitchEntranceGameserverDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchEntranceGameserverDeleteLogic {
	return &SwitchEntranceGameserverDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SwitchEntranceGameserverDeleteLogic) SwitchEntranceGameserverDelete(in *yunweiclient.DeleteSwitchEntranceGameserverReq) (*yunweiclient.SwitchEntranceGameserverCommonResp, error) {
	err := l.svcCtx.SwitchEntranceGameserverModel.DeleteSoft(l.ctx, in.Ids, in.Operation)
	if err != nil {
		return nil, xerr.NewErrMsg("批量操作开关失败，原因：" + err.Error())
	}
	//操作json文件
	for _, v := range strings.Split(in.Ids, ",") {
		one, err := l.svcCtx.SwitchEntranceGameserverModel.FindOne(l.ctx, gconv.Int64(v))
		if err != nil {
			return nil, xerr.NewErrMsg("查询单个id失败，原因：" + err.Error())
		}
		err = HandleJson(one.ConfigJsonPath, one.ConfigJsonPath, in.Operation)
		if err != nil {
			return nil, xerr.NewErrMsg("修改配置文件失败，原因：" + err.Error())
		}

	}
	return &yunweiclient.SwitchEntranceGameserverCommonResp{}, nil
}

func HandleJson(jsonFile, outFile, operation string) error {
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}
	if _, exists := result["isMaintain"]; exists {
		if operation == "stop" {
			result["isMaintain"] = true
		} else {
			result["isMaintain"] = false
		}
	}

	byteValue, err = json.Marshal(result)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	json.Indent(&out, byteValue, "", "\t")

	err = ioutil.WriteFile(outFile, out.Bytes(), 0644)
	return err
}
