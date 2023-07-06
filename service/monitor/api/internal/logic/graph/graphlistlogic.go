package graph

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/monitor/rpc/monitorclient"

	"ywadmin-v3/service/monitor/api/internal/svc"
	"ywadmin-v3/service/monitor/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GraphListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGraphListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GraphListLogic {
	return &GraphListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GraphListLogic) GraphList(req *types.ListGraphReq) (resp *types.ListGraphResp, err error) {

	var tmp monitorclient.GraphListReq
	err = copier.Copy(&tmp, req)
	if err != nil {
		return nil, xerr.NewErrMsg("复制参数失败，原因：" + err.Error())
	}
	data, err := l.svcCtx.Monitor.GraphList(l.ctx, &tmp)
	if err != nil {
		return nil, err
	}
	//fmt.Println(data.Rows)
	rows := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(data.Rows), &rows)
	if err != nil {
		return nil, xerr.NewErrMsg("转换数据失败")
	}
	resp = new(types.ListGraphResp)
	resp.Rows = rows
	return
}
