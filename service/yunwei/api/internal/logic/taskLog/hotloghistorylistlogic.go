package taskLog

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"time"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotLogHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotLogHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotLogHistoryListLogic {
	return &HotLogHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotLogHistoryListLogic) HotLogHistoryList() (resp *types.ListHotLogHistoryResp, err error) {

	ownerList, err := l.svcCtx.AdminRpc.ProjectOwnerList(l.ctx, &adminclient.ProjectOwnerReq{UserId: ctxdata.GetUidFromCtx(l.ctx)})
	if err != nil {
		return nil, err
	}
	var projectids string
	projectList := make([]*types.FilterList, 0)
	ownerTmpx := make([]string, 0)
	ownerTmpx = append(ownerTmpx, "-1")
	for _, v := range ownerList.List {
		projectList = append(projectList, &types.FilterList{
			Label: v.ProjectCn,
			Value: gconv.String(v.ProjectId),
		})
		ownerTmpx = append(ownerTmpx, gconv.String(v.ProjectId))
	}
	projectids = ctxdata.GetProjectIdsFromCtx(l.ctx)
	if projectids == "" {
		//根据用户id查询对应的项目
		projectids = strings.Join(ownerTmpx, ",")
	}

	list, err := l.svcCtx.YunWeiRpc.HotLogHistoryList(l.ctx, &yunweiclient.ListHotLogHistoryReq{
		ProjectIds: projectids,
	})

	if err != nil {
		return nil, err
	}
	// 时区
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	tmp := make([]*types.HotLogHistoryDataTree, 0)
	//获取日期数组
	tmpList := make([]string, 0)
	for _, v := range list.Rows {
		t1 := time.Unix(v.CreateTime, 0).In(Loc)
		createTimeFormat := t1.Format("2006-01-02")
		tmpList = append(tmpList, createTimeFormat)
	}
	//去重数组
	duplicate := tool.RemoveDuplicate(tmpList)
	//生成前端树形结构
	for _, v1 := range duplicate {
		tmp1 := new(types.HotLogHistoryDataTree)
		tmp2 := make([]*types.HotLogHistoryDataTree, 0)
		tmp1.Label = v1
		for _, v2 := range list.Rows {
			t1 := time.Unix(v2.CreateTime, 0).In(Loc)
			createTimeFormat := t1.Format("2006-01-02")
			if v1 == createTimeFormat {
				tmp2 = append(tmp2, &types.HotLogHistoryDataTree{
					Label: v2.HotTitle,
					Value: v2.Id,
				})
				if len(tmp2) != 0 {
					tmp1.Children = tmp2
				}
			}
		}
		tmp = append(tmp, tmp1)
	}

	//tmp := make([]*types.ListHotLogHistoryData, 0)
	//err = copier.Copy(&tmp, list.Rows)
	//if err != nil {
	//	return nil, err
	//}
	resp = new(types.ListHotLogHistoryResp)
	resp.Rows = tmp

	return
}
