package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlatformDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlatformDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformDetailLogic {
	return &PlatformDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlatformDetailLogic) PlatformDetail(in *yunweiclient.DetailPlatformReq) (*yunweiclient.DetailPlatformResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "view_platform_autoid__=", in.PlatformId)
	all, err := l.svcCtx.PlatformModel.FindListByPlatformId(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询详情失败")
	}
	if len(*all) != 1 {
		return nil, xerr.NewErrMsg("查询平台详情错误")
	}
	var tmp yunweiclient.DetailPlatformResp

	err = copier.Copy(&tmp, (*all)[0])
	if err != nil {
		return nil, xerr.NewErrMsg("复制详情失败，原因：" + err.Error())
	}

	return &tmp, nil
}
