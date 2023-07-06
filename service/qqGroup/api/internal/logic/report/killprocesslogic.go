package report

import (
	"context"
	//"syscall"

	"ywadmin-v3/service/qqGroup/api/internal/svc"
	"ywadmin-v3/service/qqGroup/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KillProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKillProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KillProcessLogic {
	return &KillProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KillProcessLogic) KillProcess(req *types.KillProcessReq) error {
	//if err := syscall.Kill(-req.KillPid, syscall.SIGKILL); err != nil {
	//	result := fmt.Sprintf("进程无法杀掉: %d, 错误信息: %s", req.KillPid, err)
	//	fmt.Println(result)
	//	return errors.New(result)
	//}
	return nil
}
