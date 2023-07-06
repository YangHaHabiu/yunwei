package captcha

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/api/internal/logic/common"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCaptchaLogic {
	return &CheckCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckCaptchaLogic) CheckCaptcha(req *types.CheckCaptchaReq) error {

	err := common.Check(l.svcCtx, req.CaptchaType, req.ClientUid, req.PointJson)
	if err != nil {
		return xerr.NewErrMsg("检测验证码失败，原因：" + err.Error())
	}
	return nil
}
