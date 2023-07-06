package msgApi

import (
	"context"
	"fmt"
	"time"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/sendMsg/api/internal/logic/common"
	"ywadmin-v3/service/sendMsg/model"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/sendMsg/api/internal/svc"
	"ywadmin-v3/service/sendMsg/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMsgApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMsgApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMsgApiLogic {
	return &ListMsgApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMsgApiLogic) ListMsgApi(req *types.ListMsgApiReq) (*types.ListMsgApiResp, error) {
	var (
		count int64

		all *[]model.SendMsgRecord
		err error
		//	start     int64
		//	end       int64
		timeRange string
	)
	list := make([]*types.ListMsgApiData, 0)
	//year, month, day := time.Now().Date()
	//location, _ := time.LoadLocation("Asia/Shanghai") // 这一步把错误忽略了，时区用Shanghai是因为没有Beijing
	//unix := time.Date(year, month, day, 0, 0, 0, 0, location).Unix()

	filters := make([]interface{}, 0)

	if req.DateRange == "" {
		nows := time.Now()
		date := nows.AddDate(0, 0, -7)
		//currentDate := nows.Format("2006-01-02")
		//beforeDate := date.Format("2006-01-02")
		currentDate := nows.Unix()
		beforeDate := date.Unix()
		timeRange = fmt.Sprintf("%d,%d", beforeDate, currentDate)
	} else {
		timeRange = req.DateRange
	}

	if req.MsgType == "" {
		if l.svcCtx.Config.ApiAuthKey.DefaultChannel == "" {
			req.MsgType = "feishu"
		} else {
			req.MsgType = l.svcCtx.Config.ApiAuthKey.DefaultChannel
		}
	}

	filters = append(filters, "msg_title__like", req.MsgTitle,
		"msg_content__like", req.MsgContent,
		"status__in", req.Status,
		"msg_type__in", req.MsgType,
		"send_type__in", req.SendType,
		"create_date__xrange", timeRange,
		"msg_level__in", req.MsgLevel,
	)

	count, _ = l.svcCtx.SendMsgRecordModel.Count(l.ctx, filters...)

	all, err = l.svcCtx.SendMsgRecordModel.FindPageListByPage(l.ctx, req.Current, req.PageSize, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}
	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	levelList, err := common.GetAlarmLevel(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	msgType, err := common.GetAlarmMsgType(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	status, err := common.GetAlarmMsgStatus(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	sendType, err := common.GetAlarmSendType(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "标题",
			Value: "msgTitle",
			Types: "input",
		},

		{
			Label: "内容",
			Value: "msgContent",
			Types: "input",
		},
		{
			Label:    "报警等级",
			Types:    "select",
			Value:    "msgLevel",
			Children: levelList,
		},
		{
			Label:    "报警渠道",
			Value:    "msgType",
			Types:    "select",
			Children: msgType,
		},

		{
			Label:    "消息类型",
			Value:    "sendType",
			Types:    "select",
			Children: sendType,
		},
		{
			Label:    "消息状态",
			Types:    "select",
			Value:    "status",
			Children: status,
		},
		{
			Label: "报警时间",
			Value: "dateRange",
			Types: "daterange",
		},
	}
	return &types.ListMsgApiResp{
		Rows:   list,
		Total:  count,
		Filter: filterList,
	}, nil
}
