package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"ywadmin-v3/common/constant"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"github.com/TestsLing/aj-captcha-go/model/vo"
	"github.com/TestsLing/aj-captcha-go/util"
	"github.com/zeromicro/go-zero/core/logx"
)

// 校验验证码
func Check(svcCtx *svc.ServiceContext, captchaType string, token string, pointJson string) error {
	codeKey := fmt.Sprintf(constant.CodeKeyPrefix, token)
	cachePointInfo, _ := svcCtx.RedisClient.Get(codeKey)
	if cachePointInfo == "" {
		return errors.New("获取验证码缓存信息失败")
	}

	if captchaType == "clickWord" {
		// 解析结构体
		var cachePoint []vo.PointVO
		var userPoint []vo.PointVO
		err := json.Unmarshal([]byte(cachePointInfo), &cachePoint)
		if err != nil {
			return err
		}
		// 解密前端传递过来的数据
		userPointJson := util.AesDecrypt(pointJson, cachePoint[0].SecretKey)
		err = json.Unmarshal([]byte(userPointJson), &userPoint)
		if err != nil {
			return err
		}

		for i, pointVO := range cachePoint {
			targetPoint := userPoint[i]
			fontsize := 18
			if targetPoint.X-fontsize > pointVO.X || targetPoint.X > pointVO.X+fontsize || targetPoint.Y-fontsize > pointVO.Y || targetPoint.Y > pointVO.Y+fontsize {
				return errors.New("文字验证失败")
			}

		}

	} else {
		// 解析结构体
		cachePoint := &vo.PointVO{}
		userPoint := &vo.PointVO{}
		err := json.Unmarshal([]byte(cachePointInfo), cachePoint)
		if err != nil {
			return err
		}
		// 解密前端传递过来的数据
		userPointJson := util.AesDecrypt(pointJson, cachePoint.SecretKey)
		err = json.Unmarshal([]byte(userPointJson), userPoint)
		if err != nil {
			return err
		}

		// 校验两个点是否符合
		if math.Abs(float64(cachePoint.X-userPoint.X)) <= float64(10) && cachePoint.Y == userPoint.Y {
			return nil
		}
		return errors.New("滑动验证失败")
	}

	return nil
}

// 根据字典类型获取字典信息
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

// 刷新策略
func FlushStrategy(svcCtx *svc.ServiceContext, ctx context.Context, userId int64, uName string) {
	//把能访问的url存在在redis，在identity中检验
	strategyList, _ := svcCtx.AdminRpc.UserStrategyList(ctx, &adminclient.UserStrategyInfoReq{
		Name: ctxdata.GetUnameFromCtx(ctx),
	})
	result := make([]string, 0)
	if strategyList.StgroupStJson != "[]" {
		var mm []types.StrategyJson
		tmp := make([]string, 0)
		//去重字典
		flagMap := make(map[string]bool, 0)
		json.Unmarshal([]byte(strategyList.StgroupStJson), &mm)

		for _, v := range mm {
			tmp = append(tmp, v.AccessUrls...)
		}
		for _, v := range tmp {
			if _, ok := flagMap[v]; ok {
				continue
			}
			flagMap[v] = true
			result = append(result, v)
		}
	}
	err := svcCtx.RedisClient.Set(fmt.Sprintf(globalkey.CacheUserAuthKey, userId), strings.Join(result, ","))
	if err != nil {
		logx.WithContext(ctx).Errorf("设置用户：%s,权限到redis异常: %+v", uName, err)
	}

}
