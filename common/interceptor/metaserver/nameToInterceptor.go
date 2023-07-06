package metaserver

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"ywadmin-v3/common/ctxdata"
)

func NameToInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.New(map[string]string{
		"uname":    ctxdata.GetUnameFromCtx(ctx),
		"uid":      gconv.String(ctxdata.GetUidFromCtx(ctx)),
		"nickName": ctxdata.GetNickNameFromCtx(ctx),
	})
	ctx = metadata.NewOutgoingContext(ctx, md)
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}
	return nil
}
