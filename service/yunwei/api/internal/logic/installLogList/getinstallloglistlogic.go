package installLogList

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInstallLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInstallLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInstallLogListLogic {
	return &GetInstallLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInstallLogListLogic) GetInstallLogList(req *types.ListInstallLogListReq) (resp *types.ListInstallLogListResp, err error) {
	list, err := l.svcCtx.YunWeiRpc.TaskGetInstallLogList(l.ctx, &yunweiclient.ListInstallLogListReq{GameName: req.GameName})
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListInstallLogListData, 0)
	err = copier.Copy(&tmp, list.Data)
	if err != nil {
		return nil, xerr.NewErrMsg("复制文件列表失败，原因：" + err.Error())
	}
	resp = new(types.ListInstallLogListResp)
	resp.Rows = tmp
	return
}
