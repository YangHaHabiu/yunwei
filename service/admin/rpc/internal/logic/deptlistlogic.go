package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListLogic) DeptList(in *adminclient.DeptListReq) (*adminclient.DeptListResp, error) {
	count, _ := l.svcCtx.DeptModel.Count(l.ctx)
	all, err := l.svcCtx.DeptModel.FindAll(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询机构列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrCode(xerr.ADMIN_DEPTSELECT_ERROR)
	}

	var list []*adminclient.DeptListData
	for _, dept := range *all {

		list = append(list, &adminclient.DeptListData{
			Id:             dept.Id,
			Name:           dept.Name,
			ParentId:       dept.ParentId,
			OrderNum:       dept.OrderNum,
			CreateBy:       dept.CreateBy,
			CreateTime:     dept.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dept.LastUpdateBy,
			LastUpdateTime: dept.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        dept.DelFlag,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询机构列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &adminclient.DeptListResp{
		Total: count,
		List:  list,
	}, nil
}
