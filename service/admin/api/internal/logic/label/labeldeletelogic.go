package label

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLabelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelDeleteLogic {
	return &LabelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LabelDeleteLogic) LabelDelete(req *types.DeleteLabelReq) (err error) {

	_, err = l.svcCtx.AdminRpc.LabelDelete(l.ctx, &admin.LabelDeleteReq{
		LabelId: req.LabelId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除异常:%s", req.LabelId, err.Error())
		return err
	}
	return
}
