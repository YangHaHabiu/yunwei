package dept

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptUpdateLogic) DeptUpdate(req *types.UpdateDeptReq) error {
	_, err := l.svcCtx.AdminRpc.DeptUpdate(l.ctx, &admin.DeptUpdateReq{
		Id:       req.Id,
		Name:     req.Name,
		OrderNum: req.OrderNum,

		LastUpdateBy: ctxdata.GetUnameFromCtx(l.ctx),
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新机构信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}

	return nil
}
