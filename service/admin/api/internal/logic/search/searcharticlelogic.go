package search

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchArticleLogic {
	return &SearchArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchArticleLogic) SearchArticle(req *types.SearchReq) (*types.SearchResp, error) {
	list, err := l.svcCtx.AdminRpc.SearchList(l.ctx, &adminclient.SearchReq{
		Keyword: req.Keyword,
	})

	if err != nil {
		return nil, err
	}

	tmp := make([]*types.ArticleReq, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制搜索信息失败，原因：" + err.Error())
	}
	return &types.SearchResp{
		Rows: tmp,
	}, nil
}
