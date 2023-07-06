package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetListLogic {
	return &AssetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssetListLogic) AssetList(in *yunweiclient.AssetListReq) (*yunweiclient.AssetListResp, error) {

	var (
		count int64
		list  []*yunweiclient.ViewAssets
		all   *[]model.ViewAssets
		err   error
	)
	filters := make([]interface{}, 0)
	filters = append(filters,
		"view_asset_id__=", in.AssetId,
		"view_outer_ip@view_inner_ip__or__regexp", in.Ips,
		"view_user_project_id__in", in.ProjectIds,
		"view_asset_ownership_company_id__in", in.OwnershipCompanyIds,
		"view_recycle_type__=", in.RecycleType,
		"view_provider_id__in", in.Provider,
		"view_clean_type__=", in.CleanType,
		"view_init_type__=", in.InitType,
		"label_names__like", in.Label,
		"view_host_role_cn__like", in.HostRoleCn,
	)

	count, _ = l.svcCtx.AssetModel.Count(l.ctx, filters...)

	if in.Current == 0 && in.PageSize == 0 {
		all, err = l.svcCtx.AssetModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.AssetModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	}

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	for _, v := range *all {

		list = append(list, &yunweiclient.ViewAssets{
			ViewAssetId:                      v.ViewAssetId.Int64,
			ViewOuterIp:                      v.ViewOuterIp.String,
			ViewInnerIp:                      v.ViewInnerIp.String,
			ViewHostRoleId:                   v.ViewHostRoleId.String,
			ViewProviderId:                   v.ViewProviderId.String,
			ViewProviderNameEn:               v.ViewProviderNameEn.String,
			ViewProviderNameCn:               v.ViewProviderNameCn.String,
			ViewHardwareInfo:                 v.ViewHardwareInfo.String,
			ViewSshPort:                      v.ViewSshPort.String,
			ViewInitType:                     v.ViewInitType.String,
			ViewCleanType:                    v.ViewCleanType.String,
			ViewInitLoginInfo:                v.ViewInitLoginInfo.String,
			ViewChangeStatusRemark:           v.ViewChangeStatusRemark.String,
			ViewRemark:                       v.ViewRemark.String,
			ViewAssetCreateTime:              v.ViewAssetCreateTime.String,
			ViewAssetUpdateTime:              v.ViewAssetUpdateTime.String,
			ViewAssetDelFlag:                 v.ViewAssetDelFlag.Int64,
			ViewPrId:                         v.ViewPrId.String,
			ViewAssetOwnershipCompanyId:      v.ViewAssetOwnershipCompanyId.Int64,
			ViewAssetOwnershipCompanyCn:      v.ViewAssetOwnershipCompanyCn.String,
			ViewAssetOwnershipCompanyEn:      v.ViewAssetOwnershipCompanyEn.String,
			ViewAssetOwnershipCompanyDeleted: v.ViewAssetOwnershipCompanyDeleted.Int64,
			ViewServerAffiliationDeleted:     v.ViewServerAffiliationDeleted.Int64,
			ViewUserCompanyId:                v.ViewUserCompanyId.Int64,
			ViewUserCompanyCn:                v.ViewUserCompanyCn.String,
			ViewUserCompanyEn:                v.ViewUserCompanyEn.String,
			ViewUserCompanyDeleted:           v.ViewUserCompanyDeleted.Int64,
			ViewUserProjectId:                v.ViewUserProjectId.Int64,
			ViewUserProjectCn:                v.ViewUserProjectCn.String,
			ViewUserProjectEn:                v.ViewUserProjectEn.String,
			ViewUserProjectDeleted:           v.ViewUserProjectDeleted.Int64,
			ViewHostRoleCn:                   v.ViewHostRoleCn.String,
			ViewRecycleType:                  v.ViewRecycleType.String,
			LabelNames:                       v.LabelNames.String,
			ViewAccelerateDomain:             v.ViewAccelerateDomain.String,
		})
	}

	return &yunweiclient.AssetListResp{
		Total: count,
		List:  list,
	}, nil
}
