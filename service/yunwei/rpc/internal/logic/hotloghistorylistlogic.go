package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"time"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotLogHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotLogHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotLogHistoryListLogic {
	return &HotLogHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HotLogHistoryListLogic) HotLogHistoryList(in *yunweiclient.ListHotLogHistoryReq) (*yunweiclient.ListHotLogHistoryResp, error) {

	var (
		list []*yunweiclient.ListHotLogHistoryData
		all  *[]model.HotLogHistory
		err  error
	)

	filters := make([]interface{}, 0)
	t := time.Now()
	beforeMonth := t.AddDate(0, 0, -15)
	format := beforeMonth.Format("2006-01-02")
	parse, _ := time.ParseInLocation("2006-01-02", format, time.Local)
	//fmt.Println(parse.Unix(), format)
	filters = append(filters, "create_time__>=", parse.Unix())
	filters = append(filters, "project_id__in", in.ProjectIds)
	all, err = l.svcCtx.HotLogHistoryModel.FindAll(l.ctx, filters...)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ListHotLogHistoryResp{
		Rows: list,
	}, nil
}
