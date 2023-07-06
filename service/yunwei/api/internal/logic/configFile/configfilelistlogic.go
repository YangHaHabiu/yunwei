package configFile

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileListLogic {
	return &ConfigFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileListLogic) ConfigFileList(req *types.ListConfigFileReq) (resp *types.ListConfigFileResp, err error) {
	tmp := make([]*types.ListConfigFileData, 0)
	list, err := l.svcCtx.YunWeiRpc.ConfigFileList(l.ctx, &yunweiclient.ListConfigFileReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	return &types.ListConfigFileResp{
		Rows:  tmp,
		Total: list.Total,
	}, nil
}
