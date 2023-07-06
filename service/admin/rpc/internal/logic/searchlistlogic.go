package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchListLogic {
	return &SearchListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// search rpc start
func (l *SearchListLogic) SearchList(in *adminclient.SearchReq) (*adminclient.SearchResp, error) {
	//在标签系统、全局搜索中显示		view_system_show[1:仅在标签中显示;2:仅在全局搜索中显示;3:标签+全局搜索中均显示]
	search, err := l.svcCtx.LabelGlobalModel.FindAllBySearch(l.ctx, "view_data_content__like", in.Keyword, "view_system_show__in", "2,3")
	if err != nil {
		return nil, xerr.NewErrMsg("全局搜索失败，原因是：" + err.Error())
	}

	var tmp []*adminclient.ArticleReq
	err = copier.Copy(&tmp, search)
	if err != nil {
		return nil, xerr.NewErrMsg("复制搜索信息失败，原因：" + err.Error())
	}
	return &adminclient.SearchResp{
		Rows: tmp,
	}, nil
}
