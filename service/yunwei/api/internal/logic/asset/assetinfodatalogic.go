package asset

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"strings"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetInfoDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetInfoDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetInfoDataLogic {
	return &AssetInfoDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetInfoDataLogic) AssetInfoData() (*types.AssetInfoDataResp, error) {

	ownerList, err := l.svcCtx.AdminRpc.ProjectOwnerList(l.ctx, &adminclient.ProjectOwnerReq{UserId: ctxdata.GetUidFromCtx(l.ctx)})
	if err != nil {
		return nil, err
	}
	ownerTmpx := make([]string, 0)
	ownerTmpx = append(ownerTmpx, "-1")
	for _, v := range ownerList.List {
		ownerTmpx = append(ownerTmpx, gconv.String(v.ProjectId))
	}

	list, err := l.svcCtx.AdminRpc.ProjectList(l.ctx, &adminclient.ProjectListReq{
		Current:    0,
		PageSize:   0,
		Status:     "-1",
		ProjectIds: strings.Join(ownerTmpx, ","),
	})
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListProjectData, 0)
	err = copier.Copy(&tmp, list.List)
	if err != nil {
		return nil, err
	}

	return &types.AssetInfoDataResp{
		ViewCompanyProjectView: tmp,
	}, nil
}
