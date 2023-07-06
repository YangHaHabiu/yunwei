package asset

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebSshTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebSshTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebSshTreeLogic {
	return &GetWebSshTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebSshTreeLogic) GetWebSshTree(req *types.GetWebSshReq) (resp *types.GetWebSshResp, err error) {
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
	if req.ProjectIds == "" {
		projectids = ctxdata.GetProjectIdsFromCtx(l.ctx)
		if projectids == "" {
			//根据用户id查询对应的项目
			projectids = strings.Join(ownerTmpx, ",")
		}
	} else {
		projectids = req.ProjectIds
	}
	list, err := l.svcCtx.YunWeiRpc.GetWebSshTree(l.ctx, &yunweiclient.GetWebSshReq{
		ProjectIds: projectids,
	})
	if err != nil {
		return nil, err
	}
	result := new(types.WebSshDataTree)
	for _, v := range list.Data {
		if v != "" {
			fmt.Println(v)
			tmp := new(types.WebSshDataTree)
			err = json.Unmarshal([]byte(v), &tmp)
			if err != nil {
				return nil, xerr.NewErrMsg("解析字典数据失败，原因：" + err.Error())
			}
			result.Key = "公司"
			result.Value = "公司"
			result.Children = append(result.Children, tmp.Children...)
		}

	}

	resp = new(types.GetWebSshResp)
	resp.Rows = result
	return
}
