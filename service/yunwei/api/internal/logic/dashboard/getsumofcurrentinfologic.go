package dashboard

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSumOfCurrentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSumOfCurrentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSumOfCurrentInfoLogic {
	return &GetSumOfCurrentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSumOfCurrentInfoLogic) GetSumOfCurrentInfo(req *types.GetSumOfCurrentInfoListReq) (resp *types.GetSumOfCurrentInfoListResp, err error) {
	//个人项目值及列表
	projectIds, _, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	list, err := l.svcCtx.YunWeiRpc.GetSumOfCurrentInfo(l.ctx, &yunweiclient.GetSumOfCurrentInfoListReq{ProjectIds: projectIds})
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.GetSumOfCurrentInfoData, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制信息出错，原因：" + err.Error())
	}
	resp = new(types.GetSumOfCurrentInfoListResp)
	resp.Rows = tmp
	return
}
