package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"google.golang.org/grpc/metadata"
	"strings"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksUpdateLogic {
	return &TasksUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TasksUpdateLogic) TasksUpdate(in *yunweiclient.UpdateTasksReq) (*yunweiclient.TasksCommonResp, error) {
	err := TaskCommon(l.svcCtx, l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.TasksModel.Update(l.ctx, in)
	if err != nil {
		return nil, xerr.NewErrMsg("更新任务失败，原因：" + err.Error())
	}

	return &yunweiclient.TasksCommonResp{}, nil
}

func TaskCommon(svcCtx *svc.ServiceContext, ctx context.Context, id int64) error {
	one, err := svcCtx.TasksModel.FindOne(ctx, id)
	if err != nil {
		return xerr.NewErrMsg("查询单条任务失败")
	}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		uid := md.Get("uid")[0]
		list, err := svcCtx.AdminRpc.UserList(ctx, &adminclient.UserListReq{UserId: gconv.Int64(uid)})
		if err != nil {
			return xerr.NewErrMsg("查询任务用户信息失败")
		}
		userObj := list.List[0]
		if userObj.Name != globalkey.SuperUserName {
			if !strings.Contains(userObj.DeptName, "运维") {
				if one.CreateBy != uid {
					return xerr.NewErrMsg("非本人操作，禁止停止")
				}
			}
		}
	}
	return nil
}
