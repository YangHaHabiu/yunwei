package insideVersion

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/api/internal/logic/common"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVersionInfoLogic {
	return &GetVersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVersionInfoLogic) GetVersionInfo(req *types.GetVersionInfoReq) (resp *types.GetVersionInfoResp, err error) {

	list, err := l.svcCtx.IntranetRpc.InsideVersionList(l.ctx, &intranetclient.ListInsideVersionReq{
		Current:   0,
		PageSize:  0,
		VersionId: req.VersionId,
	})
	if err != nil {
		return nil, err
	}
	if len(list.Rows) != 1 {
		return nil, xerr.NewErrMsg("查询版本信息失败")
	}
	byTypes, err := common.GetDictListByTypes(l.svcCtx, l.ctx, "inside_config_info", list.Rows[0].VersionType)
	if err != nil {
		return nil, err
	}
	if len(byTypes) != 1 {
		return nil, xerr.NewErrMsg("查询字典失败")
	}

	versionInfo := make([]*types.GetVersionInfoData, 0)
	if list.Rows[0].VersionType == "svn" {
		versionInfo = []*types.GetVersionInfoData{
			{"代码同步", "2022-09-29 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuronghua", "r62247"},
			{"代码同步", "2022-09-219 01:27:27 +0800 (Thu, 29 Sep 2022)", "1111", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-29 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuronghua", "r62247"},
			{"代码同步", "2022-09-219 01:27:27 +0800 (Thu, 29 Sep 2022)", "1111", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
			{"代码同步", "2022-09-292 01:27:27 +0800 (Thu, 29 Sep 2022)", "liuro333nghua", "r62247"},
		}
	} else {
		versionInfo = []*types.GetVersionInfoData{
			{"dddd", "Thu Oct 27 20:29:00 2022 +0800", "Jenkins", "94b7e3d4a733c3b2d5886cbc8c2744b8347400cc"},
			{"dddd", "Thu Oct 27 20:29:00 2022 +0800", "Jenkins", "94b7e3d4a733c3b2d5886cbc8c2744b8347400cc"},
			{"dddd", "Thu Oct 27 20:29:00 2022 +0800", "Jenkins", "94b7e3d4a733c3b2d5886cbc8c2744b8347400cc"},
		}
	}

	resp = new(types.GetVersionInfoResp)
	resp.Rows = versionInfo

	return resp, nil
}
