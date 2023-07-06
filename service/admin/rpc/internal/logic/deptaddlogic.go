package logic

import (
	"context"
	"time"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptAddLogic {
	return &DeptAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// dept rpc start
func (l *DeptAddLogic) DeptAdd(in *adminclient.DeptAddReq) (*adminclient.DeptAddResp, error) {
	_, err := l.svcCtx.DeptModel.Insert(l.ctx, &model.SysDept{
		Name:           in.Name,
		ParentId:       in.ParentId,
		OrderNum:       in.OrderNum,
		CreateBy:       in.CreateBy,
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}

	return &adminclient.DeptAddResp{}, nil
}
