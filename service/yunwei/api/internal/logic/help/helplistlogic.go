package help

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelpListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelpListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelpListLogic {
	return &HelpListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelpListLogic) HelpList(req *types.HelpListReq) (resp *types.HelpListResp, err error) {
	list, err := l.svcCtx.YunWeiRpc.HelpList(l.ctx, &yunweiclient.HelpListReq{
		GameName: req.GameName,
	})
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.HelpListData, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制帮助信息失败，原因：" + err.Error())
	}
	resp = new(types.HelpListResp)
	resp.Rows = tmp
	return
}
