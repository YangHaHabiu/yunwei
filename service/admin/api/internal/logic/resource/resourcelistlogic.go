package resource

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"strings"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceListLogic {
	return &ResourceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceListLogic) ResourceList(req *types.ListResourceReq) (*types.ListResourceResp, error) {

	ownerList, err := l.svcCtx.AdminRpc.ProjectOwnerList(l.ctx, &adminclient.ProjectOwnerReq{UserId: ctxdata.GetUidFromCtx(l.ctx)})
	if err != nil {
		return nil, err
	}

	ownerTmpx := make([]string, 0)
	ownerTmpx = append(ownerTmpx, "-1")
	for _, v := range ownerList.List {
		ownerTmpx = append(ownerTmpx, gconv.String(v.ProjectId))
	}

	list, err := l.svcCtx.AdminRpc.ResourceList(l.ctx, &adminclient.ListResourceReq{
		LabelType:  req.LabelType,
		ProjectIds: strings.Join(ownerTmpx, ","),
	})

	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListResourceList, 0)
	err = copier.Copy(&tmp, list.List)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝资源类型数据失败，原因：" + err.Error())
	}
	return &types.ListResourceResp{
		Rows: tmp,
	}, nil
}
