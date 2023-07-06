package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLabelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelDeleteLogic {
	return &LabelDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LabelDeleteLogic) LabelDelete(in *adminclient.LabelDeleteReq) (*adminclient.LabelDeleteResp, error) {
	one, err := l.svcCtx.LabelGlobalModel.FindAll(l.ctx, "lg.label_id__=", in.LabelId)
	if err != nil || len(*one) > 0 {
		return nil, xerr.NewErrMsg("请先删除存在资源对象")
	}

	err = l.svcCtx.LabelModel.DeleteSoft(l.ctx, in.LabelId)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_DELETE_ERROR)
	}

	return &adminclient.LabelDeleteResp{}, nil
}
