package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResourceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceDeleteLogic {
	return &ResourceDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResourceDeleteLogic) ResourceDelete(in *adminclient.DeleteResourceReq) (*adminclient.DeleteResourceResp, error) {
	for _, v := range in.ResourceData {
		err := l.svcCtx.LabelGlobalModel.Delete(l.ctx, v.BindingId, v.LabelId, v.ProjectId, v.ResourceEn)
		if err != nil {
			return nil, xerr.NewErrMsg("资源id：" + gconv.String(v) + "删除资源信息失败-->,失败原因：" + err.Error())
		}
	}
	return &adminclient.DeleteResourceResp{}, nil
}
