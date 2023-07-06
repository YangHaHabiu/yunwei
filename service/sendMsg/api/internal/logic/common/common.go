package common

import (
	"context"
	"fmt"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/sendMsg/api/internal/svc"
	"ywadmin-v3/service/sendMsg/api/internal/types"
)

//获取报警等级列表
func GetAlarmLevel(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "alarm_level", "报警等级")
}

//获取报警发送消息类型列表
func GetAlarmMsgType(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "send_message_type", "消息类型")
}

//获取报警发送消息状态列表
func GetAlarmMsgStatus(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "send_message_status", "消息状态")
}

//获取报警发送消息类型列表
func GetAlarmSendType(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "send_type", "发送类型")
}

//根据字典类型获取字典信息
func GetDictListByTypes(svcCtx *svc.ServiceContext, ctx context.Context, dictTypes, dictLabel string) ([]*types.FilterList, error) {
	TmpList := make([]*types.FilterList, 0)
	dictList, err := svcCtx.AdminRpc.DictList(ctx, &adminclient.DictListReq{
		Pid:      -2,
		Types:    dictTypes,
		Current:  0,
		PageSize: 0,
	})

	if err != nil {
		return nil, xerr.NewErrMsg("获取" + dictLabel + "列表失败，原因：" + err.Error())
	}
	for _, v := range dictList.List {
		if dictTypes == "feature_server_type" {
			TmpList = append(TmpList, &types.FilterList{
				Label: v.Description,
				Value: v.Label,
			})
		} else if dictTypes == "host_role_type" {
			TmpList = append(TmpList, &types.FilterList{
				Label: v.Description,
				Value: v.Description,
			})
		} else {
			var tips string
			if v.Label != v.Description {
				tips = fmt.Sprintf("%s(%s)", v.Description, v.Label)
			} else {
				tips = v.Label
			}
			TmpList = append(TmpList, &types.FilterList{
				Label: tips,
				Value: v.Value,
			})
		}

	}
	return TmpList, nil
}
