package dict

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictAddLogic {
	return &DictAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictAddLogic) DictAdd(req *types.AddDictReq) error {
	_, err := l.svcCtx.AdminRpc.DictAdd(l.ctx, &admin.DictAddReq{
		Pid:         req.Pid,
		Value:       req.Value,
		Label:       req.Label,
		Types:       req.Types,
		Description: req.Description,
		Sort:        req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加字典信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}

	return nil
}
