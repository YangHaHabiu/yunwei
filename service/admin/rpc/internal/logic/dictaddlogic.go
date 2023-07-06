package logic

import (
	"context"
	"fmt"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictAddLogic {
	return &DictAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// dict rpc start
func (l *DictAddLogic) DictAdd(in *adminclient.DictAddReq) (*adminclient.DictAddResp, error) {
	if in.Pid == -1 {
		all, err2 := l.svcCtx.DictModel.FindAll(l.ctx, "pid__=", int64(-1), "types__=", in.Types)
		fmt.Println(all)
		if err2 != nil {
			return nil, xerr.NewErrMsg("查询字典类型失败")
		}
		if len(*all) > 0 {
			return nil, xerr.NewErrMsg("新增失败，存在相同字典类型")
		}
	}

	_, err := l.svcCtx.DictModel.Insert(l.ctx, &model.SysDict{
		Value:       in.Value,
		Label:       in.Label,
		Pid:         in.Pid,
		Types:       in.Types,
		Description: in.Description,
		Sort:        in.Sort,
	})

	if err != nil {
		logx.Errorf("传入参数：%s，新增失败的原因：%+v", in, err)
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}
	return &adminclient.DictAddResp{}, nil
}
