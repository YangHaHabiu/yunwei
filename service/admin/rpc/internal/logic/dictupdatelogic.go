package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictUpdateLogic) DictUpdate(in *adminclient.DictUpdateReq) (*adminclient.DictUpdateResp, error) {
	err := l.svcCtx.DictModel.Update(l.ctx, &model.SysDict{
		Id:          in.Id,
		Value:       in.Value,
		Label:       in.Label,
		Pid:         in.Pid,
		Types:       in.Types,
		Description: in.Description,
		Sort:        in.Sort,
	})

	if err != nil {
		logx.Errorf("参数为：%s,错误信息是：%+v", in, err)
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return &adminclient.DictUpdateResp{}, nil
}
