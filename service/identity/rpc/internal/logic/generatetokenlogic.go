package logic

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/identity/rpc/internal/svc"
	"ywadmin-v3/service/identity/rpc/pb"
)

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  生成token，只针对用户服务开放访问
func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId, in.UserName, in.NickName)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", in.UserId, err)
	}

	// 存入redis.
	userTokenKey := fmt.Sprintf(globalkey.CacheUserTokenKey, in.UserId)
	err = l.svcCtx.RedisClient.Setex(userTokenKey, accessToken, int(accessExpire))
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "SetnxEx err userId:%d, err:%v", in.UserId, err)
	}

	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64, userName, nickName string) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	claims[ctxdata.CtxKeyJwtUserName] = userName
	claims[ctxdata.CtxKeyJwtNickName] = nickName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
