package taskLogHistroy

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskLogHistroyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskLogHistroyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogHistroyListLogic {
	return &TaskLogHistroyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskLogHistroyListLogic) TaskLogHistroyList(req *types.ListTaskLogHistroyReq) (resp *types.ListTaskLogHistroyResp, err error) {

	list, err := l.svcCtx.YunWeiRpc.TaskLogHistroyList(l.ctx, &yunweiclient.ListTaskLogHistroyReq{TaskId: req.TaskId})

	if err != nil {
		return nil, err
	}

	var tmp types.ListTaskLogHistroyDataJson
	err = json.Unmarshal([]byte(list.Data), &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("解析字典数据失败" + err.Error())
	}

	resp = new(types.ListTaskLogHistroyResp)
	resp.Rows = tmp
	return
}
