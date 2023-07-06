package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"net/url"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/tool"

	"ywadmin-v3/service/identity/rpc/internal/svc"
	"ywadmin-v3/service/identity/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ParseTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewParseTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParseTokenLogic {
	return &ParseTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  解析token，返回对应userid，username，nickname
func (l *ParseTokenLogic) ParseToken(in *pb.ParseTokenReq) (*pb.ParseTokenResp, error) {
	claims, err := tool.ParseToken(in.Token, in.Secret)
	if err != nil {
		return nil, err
	}
	userIdStr := claims.Claims.(jwt.MapClaims)[ctxdata.CtxKeyJwtUserId]
	nick := claims.Claims.(jwt.MapClaims)[ctxdata.CtxKeyJwtNickName].(string)
	return &pb.ParseTokenResp{
		UserId:   gconv.Int64(userIdStr),
		UserName: claims.Claims.(jwt.MapClaims)[ctxdata.CtxKeyJwtUserName].(string),
		NickName: url.QueryEscape(nick),
	}, nil
}
