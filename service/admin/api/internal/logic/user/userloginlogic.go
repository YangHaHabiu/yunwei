package user

import (
	"context"
	"fmt"
	"net/http"
	"ywadmin-v3/common/constant"
	"ywadmin-v3/common/myip"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/api/internal/logic/common"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"github.com/jinzhu/copier"

	//"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq, r *http.Request) (*types.LoginResp, error) {
	if l.svcCtx.Config.VerificationCode {
		err := common.Check(l.svcCtx, req.CaptchaType, req.ClientUid, req.PointJson)
		if err != nil {
			return nil, xerr.NewErrMsg("检测验证码失败，原因：" + err.Error())
		}
		codeKey := fmt.Sprintf(constant.CodeKeyPrefix, req.ClientUid)
		l.svcCtx.RedisClient.Del(codeKey)
	}
	loginResp, err := l.svcCtx.AdminRpc.Login(l.ctx, &admin.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResp)
	//登录日志记录
	l.svcCtx.AdminRpc.LoginLogAdd(l.ctx, &adminclient.LoginLogAddReq{
		UserName: req.Username,
		Status:   "online",
		Ip:       myip.GetCurrentIP(r),
	})
	return &resp, nil
}
