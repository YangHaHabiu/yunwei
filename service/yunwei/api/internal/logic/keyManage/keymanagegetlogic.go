package keyManage

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeyManageGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeyManageGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeyManageGetLogic {
	return &KeyManageGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeyManageGetLogic) KeyManageGet(req *types.GetKeyManageReq) (resp *types.ListKeyManageData, err error) {
	get, err := l.svcCtx.YunWeiRpc.KeyManageGet(l.ctx, &yunweiclient.GetKeyManageReq{KeyId: req.KeyId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListKeyManageData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
