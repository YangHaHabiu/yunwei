package label

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLabelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelUpdateLogic {
	return &LabelUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LabelUpdateLogic) LabelUpdate(req *types.UpdateLabelReq) (err error) {
	_, err = l.svcCtx.AdminRpc.LabelUpdate(l.ctx, &admin.LabelUpdateReq{
		LabelId:     req.LabelId,
		LabelValues: req.LabelValues,
		LabelName:   req.LabelName,
		LabelType:   req.LabelType,
		LabelRemark: req.LabelRemark,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return
}
