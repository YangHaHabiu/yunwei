package taskLogHistroy

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskLogHistroyDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskLogHistroyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogHistroyDetailLogic {
	return &TaskLogHistroyDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskLogHistroyDetailLogic) TaskLogHistroyDetail(req *types.DetailTaskLogHistroyReq) (resp *types.DetailTaskLogHistroyResp, err error) {

	detail, err := l.svcCtx.YunWeiRpc.TaskLogHistroyDetail(l.ctx, &yunweiclient.DetailTaskLogHistroyReq{TaskId: req.TaskId})
	if err != nil {
		return nil, err
	}
	tmp := new(types.DetailTaskLogHistroyResp)
	err = copier.Copy(&tmp, detail)
	if err != nil {
		return nil, xerr.NewErrMsg("复制单条数据失败" + err.Error())
	}
	return tmp, nil
}
