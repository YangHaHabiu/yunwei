package logic

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strings"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebSshTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWebSshTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebSshTreeLogic {
	return &GetWebSshTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWebSshTreeLogic) GetWebSshTree(in *yunweiclient.GetWebSshReq) (*yunweiclient.GetWebSshResp, error) {
	var uid string
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		uid = md.Get("uid")[0]
	}
	listStr := make([]string, 0)
	for _, v := range strings.Split(in.ProjectIds, ",") {
		filters := make([]interface{}, 0)
		filters = append(filters, "view_user_project_id__=", v)
		list, err := l.svcCtx.AssetModel.FindWebSshTreeList(l.ctx, uid, filters...)
		if err != nil {
			return nil, xerr.NewErrMsg("查询webssh树形结构失败，原因：" + err.Error())
		}
		if list.CompanyHost.String != "" {
			listStr = append(listStr, list.CompanyHost.String)
		}
	}
	//fmt.Println(listStr)
	return &yunweiclient.GetWebSshResp{
		Data: listStr,
	}, nil
}
