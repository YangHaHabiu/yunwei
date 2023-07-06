package label

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLabelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelAddLogic {
	return &LabelAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LabelAddLogic) LabelAdd(req *types.AddLabelReq) (err error) {
	_, err = l.svcCtx.AdminRpc.LabelAdd(l.ctx, &admin.LabelAddReq{
		LabelValues: req.LabelValues,
		LabelName:   req.LabelName,
		CreateBy:    ctxdata.GetUnameFromCtx(l.ctx),
		LabelRemark: req.LabelRemark,
		LabelType:   req.LabelType,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return
}
