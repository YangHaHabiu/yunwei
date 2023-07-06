package logic

import (
	"context"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/metadata"
	"strings"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMaintainPlanAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanAddLogic {
	return &MaintainPlanAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// KeyManage Rpc End
func (l *MaintainPlanAddLogic) MaintainPlanAdd(in *yunweiclient.AddMaintainPlanReq) (*yunweiclient.MaintainPlanCommonResp, error) {
	var tmp model.MaintainPlan
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}

	//1、检查多个集群的源ip是否一致，不一致返回错误
	info, err := l.svcCtx.MaintainPlanModel.FindAllClusterInfo(l.ctx, "view_project_id__=", in.One.ProjectId,
		"view_cluster_label_id__in", in.One.ClusterId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询项目所属集群信息失败，原因：" + err.Error())
	}

	tmpX := make([]string, 0)
	clusterIds := make([]string, 0)
	for _, v := range *info {
		outerIp := strings.TrimSpace(v.OuterIp)
		if outerIp != "" {
			tmpX = append(tmpX, outerIp)
		}
		clusterIds = append(clusterIds, gconv.String(v.LabelId))
	}

	//去重数组
	fmt.Println(tmpX)
	duplicate := tool.RemoveDuplicate(tmpX)
	clusterIds = tool.RemoveDuplicate(clusterIds)
	fmt.Println(duplicate)
	//判断去重后的列表个数
	if len(duplicate) != 1 {
		return nil, xerr.NewErrMsg("所属集群存在不一样的源IP，请检查")
	}
	var uid string
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		uid = md.Get("uid")[0]
	}
	tmp.CreateBy = uid
	tmp.ClusterId = strings.Join(clusterIds, ",")
	_, err = l.svcCtx.MaintainPlanModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &yunweiclient.MaintainPlanCommonResp{}, nil
}
