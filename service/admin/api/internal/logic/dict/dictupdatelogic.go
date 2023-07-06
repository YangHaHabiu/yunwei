package dict

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictUpdateLogic) DictUpdate(req *types.UpdateDictReq) error {
	_, err := l.svcCtx.AdminRpc.DictUpdate(l.ctx, &admin.DictUpdateReq{
		Id:          req.Id,
		Value:       req.Value,
		Pid:         req.Pid,
		Label:       req.Label,
		Types:       req.Types,
		Description: req.Description,
		Sort:        req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新字典信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}

	return nil
}
