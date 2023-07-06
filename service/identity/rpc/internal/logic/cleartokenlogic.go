package logic

import (
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xerr"
	"context"
	"github.com/pkg/errors"

	"ywadmin-v3/service/identity/rpc/internal/svc"
	"ywadmin-v3/service/identity/rpc/pb"
"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)
var ErrClearTokenError = xerr.NewErrMsg("退出token失败")

type ClearTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTokenLogic {
	return &ClearTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  清除token，只针对用户服务开放访问
func (l *ClearTokenLogic) ClearToken(in *pb.ClearTokenReq) (*pb.ClearTokenResp, error) {
	userTokenKey := fmt.Sprintf(globalkey.CacheUserTokenKey, in.UserId)
	if _, err := l.svcCtx.RedisClient.Del(userTokenKey); err != nil {
		return nil, errors.Wrapf(ErrClearTokenError, "userId:%d,err:%v", in.UserId, err)
	}

	return &pb.ClearTokenResp{
		Ok: true,
	}, nil
}
