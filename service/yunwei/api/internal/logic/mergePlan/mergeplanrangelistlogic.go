package mergePlan

import (
	"context"
	"encoding/json"
	"fmt"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanRangeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergePlanRangeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanRangeListLogic {
	return &MergePlanRangeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergePlanRangeListLogic) MergePlanRangeList(req *types.GetMergeListTreeByPlatIdAndProIdReq) (resp *types.GetMergeListTreeByPlatIdAndProIdResp, err error) {
	list, err := l.svcCtx.YunWeiRpc.MergePlanRangeList(l.ctx, &yunweiclient.GetMergeListTreeByPlatIdAndProIdReq{
		ProjectId:  req.ProjectId,
		PlatformId: req.PlatformId,
	})
	if err != nil {
		return nil, err
	}
	data := make([]*types.VueTree, 0)
	if list.Data != "" {
		err = json.Unmarshal([]byte(list.Data), &data)
		if err != nil {
			return nil, xerr.NewErrMsg(fmt.Sprintf("解析维护计划树形json失败,失败：%v", err))
		}
	}
	resp = new(types.GetMergeListTreeByPlatIdAndProIdResp)
	resp.MeRangeTreeData = data
	return
}
