package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeyManageGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKeyManageGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeyManageGetLogic {
	return &KeyManageGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KeyManageGetLogic) KeyManageGet(in *yunweiclient.GetKeyManageReq) (*yunweiclient.ListKeyManageData, error) {
	one, err := l.svcCtx.KeyManageModel.FindOne(l.ctx, in.KeyId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp yunweiclient.ListKeyManageData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
