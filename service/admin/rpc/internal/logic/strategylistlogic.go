package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StrategyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StrategyListLogic {
	return &StrategyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// strategy rpc start
func (l *StrategyListLogic) StrategyList(in *adminclient.StrategyListReq) (*adminclient.StrategyListResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters,
		"st_pid__=", in.StPid,
		"st_level__=", in.StLevel,
		"st_name__=", in.StName,
		"st_urls__in", in.StUrls,
		"st_is_auth__=", in.StIsAuth,
	)

	all, err := l.svcCtx.StrategyModel.FindAll(l.ctx, filters...)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询策略列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*adminclient.StrategyListData
	for _, info := range *all {
		list = append(list, &adminclient.StrategyListData{
			Id:       info.Id,
			StName:   info.StName,
			StLevel:  info.StLevel,
			StPid:    info.StPid,
			StUrls:   info.StUrls,
			StRemark: info.StRemark,
			StIsAuth: info.StIsAuth,
		})
	}

	return &adminclient.StrategyListResp{List: list}, nil
}
