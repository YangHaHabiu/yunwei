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

type KeyManageAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeyManageAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeyManageAddLogic {
	return &KeyManageAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeyManageAddLogic) KeyManageAdd(req *types.AddKeyManageReq) error {
	var tmp yunweiclient.KeyManageCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.KeyManageAdd(l.ctx, &yunweiclient.AddKeyManageReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil

}
