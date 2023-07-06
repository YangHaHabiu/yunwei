package stgroup

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"strings"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStgrouInfoTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStgrouInfoTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStgrouInfoTreeLogic {
	return &GetStgrouInfoTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStgrouInfoTreeLogic) GetStgrouInfoTree(req *types.GetStgrouInfoTreeReq) (resp *types.GetStgrouInfoTreeResp, err error) {

	//根据stgroupid查询策略
	list, err := l.svcCtx.AdminRpc.StgroupList(l.ctx, &adminclient.StgroupListReq{
		Current:  1,
		PageSize: 1,
		Id:       req.StgroupId,
	})
	if err != nil {
		return nil, err
	}
	if len(list.List) != 1 {
		return nil, xerr.NewErrMsg(fmt.Sprintf("id:%d查询策略信息失败", req.StgroupId))
	}

	//解析字典
	var tmp types.StrategyJson
	err = json.Unmarshal([]byte(list.List[0].StJson), &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("解析策略json失败,错误信息：" + err.Error())
	}

	//拼凑第二级和第一级条件语句
	levelTmpMap := make(map[string]bool, 0)
	levelTmpList := make([]string, 0)
	for i := 1; i <= 2; i++ {
		var level string
		for _, v := range tmp.AccessUrls {
			split := strings.Split(v, "/")
			if i == 1 {
				level = fmt.Sprintf("/%s", split[1])
			} else if i == 2 {
				level = fmt.Sprintf("/%s/%s", split[1], split[2])
			}
			if _, ok := levelTmpMap[level]; ok {
				continue
			}
			levelTmpMap[level] = true
			levelTmpList = append(levelTmpList, level)
		}
	}
	tmp.AccessUrls = append(tmp.AccessUrls, levelTmpList...)

	//根据urls查询对应策略结构，组装成树形结构
	strategyList, err := l.svcCtx.AdminRpc.StrategyList(l.ctx, &adminclient.StrategyListReq{
		StUrls: strings.Join(tmp.AccessUrls, ","),
	})

	if err != nil {
		return nil, err
	}
	//生成前端策略树形结构展示
	tmpstrategyList := make([]*types.ListstrategyData, 0)
	err = copier.Copy(&tmpstrategyList, strategyList.List)
	if err != nil {
		return nil, xerr.NewErrMsg("复制策略树形结构失败，原因：" + err.Error())
	}

	tmpStgroup := new(types.ListStgroupData)
	err = copier.Copy(&tmpStgroup, list.List[0])
	if err != nil {
		return nil, xerr.NewErrMsg("复制策略组数据失败,原因：" + err.Error())
	}

	return &types.GetStgrouInfoTreeResp{
		Rows:  tmpstrategyList,
		Datas: tmpStgroup,
	}, nil
}
